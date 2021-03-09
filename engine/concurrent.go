package engine

import (
	"log"
)

// 并发版爬虫
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

// 不希望一个接口有太多的方法，不希望传参任务太沉重
type Scheduler interface {
	ReadyNotifier // 接口组合
	Submit(Request)
	WorkerChan() chan Request // 我有一个worker，给我哪个chan，better
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// 我们希望engine的工作是很快的
func (e *ConcurrentEngine) Run(seeds ...Request) { // ...表不定参数
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//  submit 初始 request
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 获取解析的result
	for {
		result := <-out
		// 存储fetch的数据
		for _, item := range result.Iterms {
			go func() {
				e.ItemChan <- item
			}()
		}

		// 发送获取的请求给workerChan
		for _, request := range result.Requests { // (1)
			e.Scheduler.Submit(request)
		}
	}
}

// 输入chan还是不适合自己造
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() { // 这里in和out形成了循环等待
		for {
			// tell scheduler i am ready
			ready.WorkerReady(in)
			request := <-in
			log.Println("start worker")
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
			log.Println("end worker")
		}
	}()
}

// 发生了循环等待，原因
// 当给开了10个worker之后，上面的for循环卡在了Submit
// 当有worker完成之后，又没办法把result送给out
// 于是形成了循环等待
// 解决方法是破坏循环等待条件
