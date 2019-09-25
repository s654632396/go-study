package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // wg的counter自减1
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) // 设置wg的counter加1，在counter归0前, wg.Wait()会让goroutinue阻塞
		go process(i, &wg)
	}
	// 阻塞,直到 wg的counter=0
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
