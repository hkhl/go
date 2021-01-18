package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	// 启动数据存储
	itemChan, err := persist.ItemSaver("crawler")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrendEngine{    // 配置爬虫引擎
		//Scheduler:   &scheduler.SimpleScheduler{},    // 简单并发调度
		Scheduler: &scheduler.QueuedScheduler{},      // 队列并发调度
		WorkerCount: 50,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{        // 配置爬虫目标信息
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}