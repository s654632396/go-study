package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WorkNum worker numbers
const WorkNum = 4

type Task struct {
	ID       int32
	Status   string
	In       string
	Out      string
	CreateAt time.Time
}

func (t *Task) run() {
	sleep := rand.Intn(1000)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	t.Status = "Completed"
}

func main() {
	var wg sync.WaitGroup
	worknum := WorkNum
	wg.Add(worknum)

	var ch chan *Task = make(chan *Task, 10)
	startWorkers(worknum, wg, ch)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 30; i++ {
		ch <- &Task{
			ID:       int32(i),
			Status:   "undo",
			In:       string(rand.Intn(10000)),
			Out:      "",
			CreateAt: time.Now(),
		}
	}
	close(ch)

	fmt.Println("all task sent.")
	wg.Wait()
	fmt.Println("all task done.")
}

func startWorkers(num int, wg sync.WaitGroup, ch chan *Task) {
	for i := 0; i <= num; i++ {
		fmt.Println("start worker :", i+1)
		go func(wg sync.WaitGroup, ch chan *Task) {
			defer wg.Done()
			//
			for t := range ch {
				// comsumer
				// do some thing
				fmt.Printf("recv task(%d). \n", t.ID)
				t.run()
			}
		}(wg, ch)
	}
}
