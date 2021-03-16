package parser

import (
	"fmt"
	"github.com/Winszheng/crawler/engine"
	"regexp"
	"strings"
)

// 解析单个城市第一页的用户链接
var (
	userRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)
	// to get:xx城市百姓/军人/公务员...征婚
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^/]+/[a-z]+)">([^<]+)</a>`)
)

// 解析城市页面的用户列表
func ParseCity(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}
	matches := userRe.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		result.Iterms = append(result.Iterms, engine.Item{
			Url:      "",
			Id:       "",
			Playload: nil,
		})
		temp := strings.Replace(string(match[1]), "album", "m", -1)
		result.Requests = append(result.Requests, engine.Request{
			Url:       temp,
			ParseFunc: ParseProfile,
		})
		fmt.Println("get user", string(match[2]), "'s", "url:", temp)
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]), //string(m[1])
			ParseFunc: ParseCity,
		})
		fmt.Println(string(m[2]), string(m[1]))
	}
	return result
}
