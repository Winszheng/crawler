package parser

import (
	"fmt"
	"github.com/Winszheng/crawler/af/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}	// 构造返回结果
	for k, match := range matches {
		if k > 3 {
			break // 为了省事，实际不应当有限制
		}

		result.Items = append(result.Items, match[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(match[1]),
			Parser: ParseCity,
		})

		fmt.Println("Parsing city:", string(match[2]))

	}
	return result
}
