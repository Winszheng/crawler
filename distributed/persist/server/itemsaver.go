package main

import (
	"flag"
	"fmt"
	"github.com/Winszheng/crawler/distributed/config"
	"github.com/Winszheng/crawler/distributed/persist"
	"github.com/Winszheng/crawler/distributed/rpcsupport"
	"github.com/olivere/elastic/v7"
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
	// 挂了就强制退出
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	// 因为ItemSaverService是用指针实现相应方法的，形参是接口，所以要传指针
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
