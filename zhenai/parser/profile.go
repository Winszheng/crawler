package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Winszheng/crowler/engine"
	"github.com/Winszheng/crowler/model"
	"log"
	"regexp"
)

// nickname
var nameRe = regexp.MustCompile(`<div class="name" data-v-44b88fba><span data-v-44b88fba>([^<]+)</span>`)

// 内心独白
var contentRe = regexp.MustCompile(`<span data-v-37f6ec3b>内心独白</span></div> <div class="content" data-v-37f6ec3b><p data-v-37f6ec3b>([^<]+)</p>`)

func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}
	if text := nameRe.FindSubmatch(contents); len(text) > 0 {
		profile.Nickname = string(text[1])
	}
	if text := contentRe.FindSubmatch(contents); len(text) > 0 {
		profile.Content = string(text[1])
	}
	// profile.Nickname = string(nameRe.FindSubmatch(contents)[1])
	// profile.Content = string(contentRe.FindSubmatch(contents)[1])

	// 个人资料和择偶条件不适合用re
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		log.Fatalln(err)
		return engine.ParseResult{}
	}
	dom.Find(".basicInfo-section>.tag").Each(func(i int, selection *goquery.Selection) {
		profile.BasicInfo = append(profile.BasicInfo, selection.Text())
	})
	dom.Find(".detailInfo-section>.tag").Each(func(i int, selection *goquery.Selection) {
		profile.DetailInfo = append(profile.DetailInfo, selection.Text())
	})
	dom.Find(".objectInfo-section>.tag").Each(func(i int, selection *goquery.Selection) {
		profile.Selection = append(profile.Selection, selection.Text())
	})
	fmt.Println("get user info:", profile)
	result := engine.ParseResult{
		Iterms: []interface{}{profile}, // 没有新的请求
	}
	return result
}
