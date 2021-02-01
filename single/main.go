package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 单任务爬虫：
// 获取并打印所有城市第一页用户的详细信息

func main() {
	printCityList(all)
}

func printCityList(contents []byte) {
	// 匹配各个城市的链接
	// []byte就`相当于`string，只是打印的时候要注意格式
	// [][]byte相当于一个切片，里面的元素都是string
	// 用正则表达式匹配城市标签，挖出城市链接和名字--()
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-1573aa7c>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for i, m := range matches {
		fmt.Printf("%d %s %s\n", i+1, m[1], m[2])
	}

}
