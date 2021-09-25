package engine

import (
	"github.com/Winszheng/crawler/af/fetcher"
	"log"
)

// 初始版engine
// 目前的操作是，不做调度，顺序执行
func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetcher(r.Url)
		if err != nil {
			log.Println("Fetching Error:", r.Url)
		}
		result := r.Parser(body, r.Url)

		requests = append(requests, result.Requests...)
	}
}
