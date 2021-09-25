package engine

import (
	"github.com/Winszheng/crawler/single/fetcher"
	"log"
)

// worker是引擎公用的
// 把Worker包装成rpc服务
func Worker(r Request) (ParseResult, error) {
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch: error fetching %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(content, r.Url), nil
}
