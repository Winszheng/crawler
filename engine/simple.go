package engine

import (
	"fmt"
	// "github.com/Winszheng/crowler/fetcher"
	// "log"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	// Run需要维护一个请求队列
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 { // 只要请求队列中还有请求，就一直处理
		// get the first request from the queue
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)

		if err != nil {
			fmt.Println(err)
			continue
		}
		requests = append(requests, result.Requests...)
	}
}
