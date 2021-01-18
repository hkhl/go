package engine

// 并发引擎
type ConcurrendEngine struct {
	Scheduler   Scheduler    // 任务调度器
	WorkerCount int          // 任务并发数量
	ItemChan    chan Item 	 // 数据保存 channel
}

// 任务调度器
type Scheduler interface {
	ReadyNotifier
	Submit(request Request) // 提交任务
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrendEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 创建 goruntine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// engine把请求任务提交给 Scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	for {
		// 接受 Worker 的解析结果
		result := <-out
		for _, item := range result.Items {
			//抓去一组数据后，进行保存
			go func() {
				e.ItemChan <- item
			}()

		}

		// 然后把 Worker 解析出的 Request 送给 Scheduler
		for _, request := range result.Requests {
			// 如果重复，则不提交任务
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}


// 创建任务，调用worker，分发goroutine
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	// 为每一个Worker创建一个channel
	go func() {
		for {
			ready.WorkerReady(in) // 告诉调度器任务空闲
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// 存放已经获取的所有Url
var visitedUrls = make(map[string]bool)

// 判断Url是否重复
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	} else {
		return false
	}
}
