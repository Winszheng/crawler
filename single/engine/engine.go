package engine

import (
	"github.com/Winszheng/crowler/single/fetcher"
	"log"
)

// 每个Request就是一个种子
// 虽然不直到为什么要叫做种子！
func Run(seeds ...Request) {
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
		for _, item := range ParseResult.Items {
			log.Printf("Got item %v ", item)
		}

	}
}