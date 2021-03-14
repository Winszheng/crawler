package engine

import (
	"github.com/Winszheng/crowler/fetcher"
	"log"
)

// worker是引擎公用的
func worker(r Request) (ParseResult, error) {
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch: error fetching %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(content, r.Url), nil
}
