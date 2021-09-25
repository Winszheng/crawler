package main

import (
	"github.com/Winszheng/crawler/distributed/config"
	"github.com/Winszheng/crawler/distributed/rpcsupport"
	"github.com/Winszheng/crawler/single/engine"
	"github.com/Winszheng/crawler/single/model"
	"testing"
	"time"
)

// TestItemSaver起了一个server，测试本身扮演client
// 验证server有没有做对
// 原先写的爬虫相当于client
func TestItemSaver(t *testing.T) {
	const host = ":1234" // 测试独有的host，因为可以根据自己测试的需要改
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second * 1)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// call save
	item := engine.Item{
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

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; %s", result, err)
	}

}
