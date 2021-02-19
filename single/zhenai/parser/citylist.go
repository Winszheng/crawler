package parser

import (
	"github.com/Winszheng/crowler/single/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`

// ParseCityList
// 这个函数需要测试
func ParseCityList(contents []byte) engine.ParseResult{
	// 匹配各个城市的链接
	// []byte就`相当于`string，只是打印的时候要注意格式
	// [][]byte相当于一个切片，里面的元素都是string
	// 用正则表达式匹配城市标签，挖出城市链接和名字==>()
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}  // 初始化要返回的结果

	// 为了省事，只爬10个城市
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),   // change []byte to string
			ParserFunc: ParseCity,
		})

		limit--
		if limit == 0 {
			break
		}
	}
	return result
}