package parser

import (
	"fmt"
	"github.com/Winszheng/crawler/af/engine"
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
	result := engine.ParseResult{}	// 构造返回结果
	matches := userRe.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		result.Items = append(result.Items, "")
		temp := strings.Replace(string(match[1]), "album", "m", -1)	// 改变请求url
		result.Requests = append(result.Requests, engine.Request{
			Url:    temp,
			Parser: ParseProfile, // 我没有用name
		})
		fmt.Println("get user", string(match[2]), "'s", "url:", temp)
	}
	return result
}