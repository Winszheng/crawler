package main

import (
	"github.com/Winszheng/crowler/single/engine"
	"github.com/Winszheng/crowler/single/zhenai/parser"
)

// 单任务爬虫：
// 获取并打印所有城市第一页用户的详细信息

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

