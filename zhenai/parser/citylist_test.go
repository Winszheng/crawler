package parser

import (
	"io/ioutil"
	"testing"
)

// 选取部分做测试
func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents, "")

	const resultSize = 470
	// const不能定义数组切片
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectedCities = []string{"阿坝", "阿克苏", "阿拉善盟"}

	if len(result.Requests) != resultSize {
		t.Errorf("expect %d request, but get %d request", resultSize, len(result.Requests))
	}

	for i := 0; i < len(expectedUrls); i++ {
		if expectedUrls[i] != result.Requests[i].Url {
			t.Errorf("expect url \"%s\", but get url \"%s\"", expectedUrls[i], result.Requests[i].Url)
		}

		if result.Iterms[i].Playload != expectedCities[i] {
			t.Errorf("expect city \"%s\", but get city \"%s\"", expectedCities[i], result.Iterms[i])
		}
	}
}
