package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go readChannel(ch, 1)
	go readChannel(ch, 2)
	go readChannel(ch, 3)

	var dataMap []string = []string{"123", "abc", "def", "mss", "ras", "dvx", "dfd", "qwe"}
	for _, data := range dataMap {
		time.Sleep(1 * time.Second)
		ch <- data
	}
	close(ch)

	time.Sleep(5 * time.Second)
	println("end.")
}

func readChannel(ch <-chan string, id int) {
	fmt.Printf("child[%d] start. \n", id)
	for {
		data, closed := <-ch
		if !closed {
			fmt.Println("channel is closed, now break down this goroutinue.")
			break
		}
		fmt.Printf("[id:%d] read data from parent: %s \n", id, data)
		// time.Sleep(1 * time.Second)
	}
	fmt.Printf("child[%d] closed. \n", id)
}
