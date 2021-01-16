package main

import (
	"fmt"
	"sync"
)

func Producer(ids chan int) {
	for id := 1; id <= 10000; id++ {
		ids <- id
	}
	close(ids)
}

func Consumer(ids chan int, workNum int)  {
	var wg sync.WaitGroup
	for i:=0; i<workNum;i++ {
		wg.Add(1)
		go func(workID int) {
			defer wg.Done()
			for id := range ids {
				fmt.Printf("id=%d already consumed by workID[%d].\n", id, workID)
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	c := make(chan int, 10)

	go Producer(c)
	Consumer(c, 10)


}
