package parser

import (
	"fmt"
	"github.com/Winszheng/crawler/distributed/config"
	"github.com/Winszheng/crawler/single/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`

func ParserCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for k, match := range matches {
		if k > 100 {
			break
		}

		result.Iterms = append(result.Iterms, engine.Item{
			Playload: match[2],
		})
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(match[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity), // 暂且
		})
		// 打印顺序 名字 url
		fmt.Printf("%d %s %s\n", k+1, match[2], match[1])
	}
	return result
}
