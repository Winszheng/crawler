package parser

import (
	"fmt"
	"github.com/Winszheng/crowler/single/engine"
	"strings"
)
import "github.com/PuerkitoBio/goquery"


func ParseProfile(contents []byte) engine.ParseResult {
//	与正则表达式相比
//	使用css选择器更适合具体的文档
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}
	// 内心独白
	dom.Find(".m-des").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	// 个人资料
	// 兴趣爱好
	// 择偶条件
	return engine.ParseResult{} // 暂且返回一个空结构体
}