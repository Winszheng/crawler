package parser

import (
	"fmt"
	"github.com/Winszheng/crowler/engine"
	"regexp"
	"strings"
)

// 解析单个城市第一页的用户链接
var (
	userRe    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	matches := userRe.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		result.Iterms = append(result.Iterms, string(match[2]))
		temp := strings.Replace(string(match[1]), "album", "m", -1)
		result.Requests = append(result.Requests, engine.Request{
			Url:       temp,
			ParseFunc: ParseProfile,
		})
		fmt.Println("get user", string(match[2]), "'s", "url:", temp)
	}

	cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
