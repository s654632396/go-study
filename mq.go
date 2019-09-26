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
	// simulate dealing task
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	// change task status to completed
	t.Status = "Completed"
}

func main() {
	var wg sync.WaitGroup
	// channel of task Queue
	var ch chan *Task = make(chan *Task, 10)
	startWorkers(WorkNum, &wg, ch)

	rand.Seed(time.Now().UnixNano())

	// producer
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

func startWorkers(num int, wg *sync.WaitGroup, ch chan *Task) {
	for i := 1; i <= num; i++ {
		wg.Add(1)
		fmt.Println("start worker :", i)
		go func(wg *sync.WaitGroup, ch chan *Task, workId int16) {
			defer wg.Done()
			//
			for t := range ch {
				t.Status = "prepare"
				// comsumer
				// do some thing
				fmt.Printf("recv task(%d). \n", t.ID)
				t.run()
				fmt.Printf("done task(%d). \n", t.ID)
			}
		}(wg, ch, int16(i))
	}
}
