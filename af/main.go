package main

import (
	"github.com/Winszheng/crawler/af/engine"
	"github.com/Winszheng/crawler/af/parser"
)

// 先把单任务版跑通，再改成并发版，在弄成有存储和前端的，再改成分布式版

func main()  {
	// 1.启动单任务版本
	// engine.Run(engine.Request{
	// 	Url:    "https://www.zhenai.com/zhenghun",
	// 	Parser: parser.ParseCityList,
	// })

	// 2.启动并发版爬虫
	e := engine.ConcurrentEngine{
		WorkerCount: 30,
	}

	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: parser.ParseCityList,
	})



}
