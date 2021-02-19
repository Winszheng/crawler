package parser

import (
	"github.com/Winszheng/crowler/single/engine"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		// http://m.zhenai.com/u/1275335590
		// https://album.zhenai.com/u/1275335590
		url := strings.Replace(string(m[1]), "album", "m", -1)
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseProfile,
		})
		println("got item user: ", string(m[2])," ", url)
	}
	return result
}
