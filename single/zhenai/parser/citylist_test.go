package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch(
	//	"http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	// verify result
	result := ParseCityList(contents)
	const resultSize = 470
	// 不能用const定义数字切片
	// 取部分出来测试
	var expectUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d",
			resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d",
			resultSize, len(result.Items))
	}

	// 查看取到的数据是否是正确的
	for i, url := range expectUrls {
		if url != result.Requests[i].Url {
			t.Errorf("expected url #%d: %s, but get %s", i,
				url, result.Requests[i].Url)
		}
	}

	for i, city := range expectedCities {
		if city != result.Items[i].(string) {   // 类型断言
			t.Errorf("expected city #%d: %s, but get %s", i,
				city, result.Items[i])
		}
	}


}
