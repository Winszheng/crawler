package engine

import (
	"context"
	"fmt"
	"github.com/Winszheng/crawler/af/fetcher"
	"github.com/Winszheng/crawler/af/model"
	"github.com/olivere/elastic/v7"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ConcurrentEngine struct {
	WorkerCount int
	In chan Request		// 用一个带缓冲的channel表征请求队列，缓冲区大小自定义
	ItemChan chan interface{}	// 存放item
}

func(e *ConcurrentEngine) Run(seeds ...Request)  {
	e.In = make(chan Request, e.WorkerCount)
	e.ItemChan = make(chan interface{}, e.WorkerCount)

	wg.Add(e.WorkerCount+1)

	go func() {
		for _, r := range seeds {
			e.In <- r
		}
	}()

	for i:=0; i<e.WorkerCount; i++ {
		go e.Worker(i)
	}

	go e.ItemSaver()
	wg.Wait()

	fmt.Println("爬虫程序执行完毕")
}

// 查询结果：http://localhost:9200/dating_profile/_search
func Save(client *elastic.Client, item model.Profile) error {
	indexService := client.Index().
		Index("dating_profile").
		Id(item.Id).
		BodyJson(item)
	resp, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("存储情况：", resp)
	return err
}

func(e *ConcurrentEngine) ItemSaver()  {
	i := 0
	client, err := elastic.NewClient(
		// 这是用来维护集群的，因为项目的集群不在本机，而在docker，所以设置成false
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err) // 连不上时，不存了，退出
	}
	for{
		select {
		case item := <-e.ItemChan:
			if it, ok := item.(model.Profile); ok {
				if err = Save(client, it); err != nil {
					log.Printf("Item Saver: error saving item %v: %v", it, err)
				}
				i++;
			}
		case <-time.After(time.Second * 100):
			fmt.Println("ItemSaver has finished")
			wg.Done()
			return
		}
	}
}

// Worker 希望并发的速度是很快的, 所以要把任务快速地分发下去
func(e *ConcurrentEngine) Worker(i int) {
	fmt.Println("Worker", i, "is running")
	for {
		select {
		case r := <-e.In:
			body, err := fetcher.Fetcher(r.Url)
			if err != nil {
				log.Println("Error at Worker:", err)
			}
			result := r.Parser(body, r.Url)
			for i, _ := range result.Requests {	// 避免每次操作的都是同一个请求
				go func(i int) {e.In <- result.Requests[i]}(i)
			}
			for i, _ := range result.Items {	// 避免每次操作的都是同一个item
				go func(i int) {e.ItemChan <- result.Items[i]}(i)
			}
		case <-time.After(time.Second*10):	// 连续10秒没有收到新请求，证明这个协程该结束了
			fmt.Println("Worker", i, "has finished")
			wg.Done()
			return
		}
	}
}