package main

import (
	"github.com/Winszheng/crawler/single/engine"
	"github.com/Winszheng/crawler/single/persist"
	"github.com/Winszheng/crawler/single/scheduler"
	"github.com/Winszheng/crawler/single/zhenai/parser"
)

func main() {
	itemsaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err) // itemSaver开不起来，干脆就不运行程序
	}
	e := &engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{}, // 因为是指针接收者所以要取地址，见下面的报错
		WorkerCount:      30,
		ItemChan:         itemsaver, // 如何存储爬到的数据
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParserCityList, "ParserCityList"), // 成员变量是个函数
	})
}

// 报错：cannot call pointer method on engine.ConcurrentEngine literal(字面的)
// 这是可不可以寻址的问题，不可以对一个字面值寻址，但是可以对变量或指针寻址，
// 所以无论是e：= engine.ConcurrentEngine{...}还是e：= &engine.ConcurrentEngine{...}都是对的
// golang会根据接收者定义帮你调整
// 如果一开始接收者是值，也就不用考虑寻址的问题了。
