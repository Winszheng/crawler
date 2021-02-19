package engine

import (
	"github.com/Winszheng/crowler/single/fetcher"
	"log"
)

// 引擎， 或者说，启动器
// 每个Request就是一个种子
func Run(seeds ...Request) {
	// 维护一个任务队列，不断调用相应的解析器
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 当任务队列>0的时候，就可以运行
	for len(requests) > 0 {
		r := requests[0]    // 把队列第一个请求拿出来
		requests = requests[1:] // 切片的剪切

		log.Printf("Fetching %s", r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue    // 处理下一个Request
		}

		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Requests...)   // 把切片打散拼接

	}
}