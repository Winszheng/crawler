package main

import (
	"flag"
	"fmt"
	"github.com/Winszheng/crawler/distributed/config"
	itemsaver "github.com/Winszheng/crawler/distributed/persist/client"
	"github.com/Winszheng/crawler/distributed/rpcsupport"
	worker "github.com/Winszheng/crawler/distributed/worker/client"
	"github.com/Winszheng/crawler/single/engine"
	"github.com/Winszheng/crawler/single/scheduler"
	"github.com/Winszheng/crawler/single/zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver_host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

// 分布式爬虫
// 爬虫服务本身相当于client，其他worker和saver服务都相当于server
// client把各种任务发到不同的server处理
// go run main.go -itemsaver_host=":1234" -worker_hosts=":9000,:9001,:9002"
func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf("%s", *itemSaverHost))
	if err != nil {
		panic(err) // itemSaver开不起来，干脆就不运行程序
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	if err != nil {
		panic(err)
	}
	e := &engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{}, // 因为是指针接收者所以要取地址，见下面的报错
		WorkerCount:      50,
		ItemChan:         itemChan, // 如何存储爬到的数据
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParserCityList, config.ParserCityList), // 成员变量是个函数
	})
}

// 使用连接池链接爬虫集群
// 通过消息传递共享数据
func createClientPool(hosts []string) chan *rpc.Client {
	// 针对传进来的hosts建client
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
