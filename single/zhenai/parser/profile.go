package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Winszheng/crowler/model"
	"github.com/Winszheng/crowler/single/engine"
	"math/rand"
	"strings"
	"time"
)


// 个人主页具体信息
func ParseProfile(contents []byte) engine.ParseResult {
	time.Sleep(time.Millisecond*time.Duration(rand.Intn(100)))

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		return engine.ParseResult{}
	}
	profile := model.Profile{}
	dom.Find(".name").Each(func(i int, selection *goquery.Selection) {
		profile.Nickname = selection.Text()
	})
	dom.Find(".content").Each(func(i int, selection *goquery.Selection) {
		profile.Des = selection.Text()
	})
	dom.Find(".basicInfo-section>.tag").Each(func(i int, selection *goquery.Selection) {
		profile.BasicInfo = append(profile.BasicInfo, selection.Text())
	})
	dom.Find(".detailInfo-section>.tag").Each(func(i int, selection *goquery.Selection) {
		profile.Detail = append(profile.Detail, selection.Text())
	})
	dom.Find(".objectInfo-section>.tag").Each(func(i int, selection *goquery.Selection) {
		profile.Selection = append(profile.Selection, selection.Text())
	})


	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{profile},
	}

	fmt.Println("got item user: ", profile.Nickname, "'s detail information: ", result.Items)
	return result
}