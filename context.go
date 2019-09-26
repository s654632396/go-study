package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	simpleCaseTwo()

}

func simpleCaseTwo() {
	var ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go watch(ctx, "a-01")
	go watch(ctx, "a-02")
	go watch(ctx, "a-03")

	time.Sleep(time.Second * 10)
	cancel()
	time.Sleep(time.Second * 3)
}

func watch(ctx context.Context, id string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("watcher[%s] is stop.\n", id)
			return
		default:
			fmt.Printf("watcher[%s] is running.\n", id)
			time.Sleep(time.Second * 1)
		}
	}
}

// simple control goroutinue stop with Context
func simpleCaseOne() {

	var ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine is go to stop.")
				return
			default:
				fmt.Println("goroutine is running.")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
