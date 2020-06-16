package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {

	var waitChan chan int
	waitChan = make(chan int, 10)

	for i := int(0); i < 10; i++ {
		go func(i int) {
			log.Printf("start goroutine %d .\n", i)
			time.Sleep(time.Duration(2 * time.Second))
			defer func() {
				log.Printf("goroutinue %d exit.\n", i)
				waitChan <- i
			}()
			runtime.Goexit() // 使用runtime.Goexit 来终止goroutine
			fmt.Print("这段不会被执行到.")
		}(i)
	}
	counter := 0
	for ch := range waitChan {
		fmt.Println("read from channel: ", ch)
		counter++
		if counter >= 10 {
			close(waitChan)
		}
	}

	log.Println("all done.")

}
