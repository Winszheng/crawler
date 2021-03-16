package parser

import (
	"fmt"
	"github.com/Winszheng/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`

func ParserCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for k, match := range matches {

		result.Iterms = append(result.Iterms, engine.Item{
			Playload: match[2],
		})
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(match[1]),
			ParseFunc: ParseCity, // 暂且
		})
		// 打印顺序 名字 url
		fmt.Printf("%d %s %s\n", k+1, match[2], match[1])
	}
	return result
}
