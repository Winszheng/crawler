package client

import (
	"github.com/Winszheng/crawler/distributed/config"
	"github.com/Winszheng/crawler/distributed/rpcsupport"
	"github.com/Winszheng/crawler/single/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() { // 持久化的逻辑
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			// ev time received a item
			// call rpc to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil { // 重试/放弃，这里选abandon
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
