package main

import (
	"flag"
	"fmt"
	"github.com/Winszheng/crawler/distributed/rpcsupport"
	"github.com/Winszheng/crawler/distributed/worker"
	"log"
)

// 参数名 默认值 帮助文档
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse() // 解析命令行参数值
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlerService{}))
}
