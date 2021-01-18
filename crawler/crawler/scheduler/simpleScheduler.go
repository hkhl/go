package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	// 此时所有 worker 共用同一个 channel，直接返回即可
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	// 创建出 workchannel
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// 为每一个Request创建goroutine
	go func() {
		s.workerChan <- request
	}()
}

//// 把初始请求发送给 Scheduler
//func (s *SimpleScheduler) ConfigMasterWorkerChan(in chan engine.Request) {
//	s.workerChan = in
//}
