package persist

import (
	"context"
	"encoding/json"
	"github.com/Winszheng/crawler/single/engine"
	"github.com/Winszheng/crawler/single/model"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/108906739",
		Id:  "108906739",
		Playload: model.Profile{
			Nickname:   "清岚",
			Content:    "我性格开朗，希望他也一样，真诚相亲，非诚勿扰",
			BasicInfo:  []string{"离异", "43岁", "天秤座(09.23-10.22)", "157cm", "工作地:上海长宁区", "月收入:5-8千", "客服专员", "中专"},
			DetailInfo: []string{"籍贯:安徽宣城", "体型:丰满", "不吸烟", "不喝酒", "住在单位宿舍", "未买车", "有孩子但不在身边", "是否想要孩子:视情况而定"},
			Selection:  []string{"41-65岁", "工作地:上海", "月薪:1.2万以上"},
		},
	}

	const index = "dating_test"
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_test").Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)
	expextedJson, err := json.Marshal(expected)
	if err != nil {
		panic(err)
	}

	if string(expextedJson) != string(resp.Source) {
		t.Errorf("got %v; expected %v", string(resp.Source), string(expextedJson))
	}

}
