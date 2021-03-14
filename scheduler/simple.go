package scheduler

import "github.com/Winszheng/crowler/engine"

// 希望simple和queued的代码能有比较好的通用性
type SimpleScheduler struct {
	wokerChan chan engine.Request
}

// simple scheduler的workerChan是共用的
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.wokerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.wokerChan = make(chan engine.Request)
	// 简单版并发，所有worker共用一个input chan
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	// 被调用了很多次，每次都是不一样的request
	go func() {
		s.wokerChan <- r
	}()
}
