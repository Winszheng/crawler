package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() { // 持久化的逻辑放在这里
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out // 这玩意return了干嘛？
}
