package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)

	go setData(ch)
	go getData(ch)

	// go setData(ch)
	// getData(ch)

	time.Sleep(1 * time.Second)
	close(ch)
	fmt.Println("end...")

}

func setData(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
}

func getData(ch chan int) {
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Println("channel is closed.")
				return
			}
			fmt.Println(data)

		default:
		}
	}
}
