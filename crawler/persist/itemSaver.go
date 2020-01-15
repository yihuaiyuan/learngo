package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("Item saver got item #%d : %v", itemCount, item)
			itemCount++
		}
	}()

	return out
}