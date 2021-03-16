package scheduler

import "github.com/Winszheng/crowler/engine"

// 队列调度器
// request和worker都是队列
// 只要既有request在排队，又有worker在排队，就active，当真的把request送给worker，再从双方队列移除
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request // 每个worker有自己的chan
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	// 希望每个chan有自己的request
	return make(chan engine.Request)
}

func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}
func (s *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	// 因为是生成workerChan和requestChan，改变了接收者，所以要使用指针接收者

	go func() {
		// 建立两个队列
		// 调度逻辑：如果既有request在排队，又有worker在排队，那么就可以把request发给worker
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0] // 不需要从队列移走，那么要怎么处理呢？
			}
			select { // 有select就尽量把chan操作放到select
			case r := <-s.requestChan: // 收到一个request就让它排队，下面同理
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: // 收到一个worker就让它排队
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: // 把request发给worker的时候，才真的移走
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
				// 	然后的问题是，如何生成这两个chan
			}
		}
	}()
}
