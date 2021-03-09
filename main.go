package main

import (
	"github.com/Winszheng/crowler/engine"
	"github.com/Winszheng/crowler/persist"
	"github.com/Winszheng/crowler/scheduler"
	"github.com/Winszheng/crowler/zhenai/parser"
)

func main() {
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, // 因为是指针接收者所以要取地址，见下面的报错
		WorkerCount: 30,
		ItemChan:    persist.ItemSaver(), // 如何存储爬到的数据
	}
	e.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParserCityList, // 成员变量是个函数
	})
}

// 报错：cannot call pointer method on engine.ConcurrentEngine literal(字面的)
// 这是可不可以寻址的问题，不可以对一个字面值寻址，但是可以对变量或指针寻址，
// 所以无论是e：= engine.ConcurrentEngine{...}还是e：= &engine.ConcurrentEngine{...}都是对的
// golang会根据接收者定义帮你调整
// 如果一开始接收者是值，也就不用考虑寻址的问题了。

// 2.需要考虑到一个问题叫方法集