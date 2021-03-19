package persist

import (
	"github.com/Winszheng/crawler/single/engine"
	persist "github.com/Winszheng/crawler/single/persist"
	"github.com/olivere/elastic/v7"
	"log"
)

// rpc.go define rpc service
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// 返回操作结果语
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.\n", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
