package persist

import (
	"context"
	"errors"
	"github.com/Winszheng/crawler/single/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

// ItemSaver
// index: 外界决定存到哪个index
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// 这是用来维护集群的，因为项目的集群不在本机，而在docker，所以设置成false
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err // 连不上时，由外面决定如何处理
	}
	out := make(chan engine.Item)
	go func() { // 持久化的逻辑
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			if err := Save(client, index, item); err != nil { // 重试/放弃，这里选abandon
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// 只存用户的个人信息
func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Id == "" {
		return errors.New("触发反爬机制，停止爬取")
	}

	indexService := client.Index().
		Index(index).
		Id(item.Id).
		BodyJson(item)

	resp, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	if item.Id == "" {
		item.Id = resp.Id
	}

	return err
}
