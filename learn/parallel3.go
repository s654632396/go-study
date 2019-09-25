package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type InnerList []int
type OuterList []InnerList

func main() {
	// setting GOMAXPROCS
	var numCPU = runtime.NumCPU()
	fmt.Printf("system usable cpu number is %d \n", numCPU)
	runtime.GOMAXPROCS(numCPU)

	var list OuterList = genList()
	// preview list
	fmt.Println("list ------------------------------------")
	fmt.Printf("%#v \n", list)
	fmt.Println("-----------------------------------------")

	// ===== 使用带缓冲区的channel，而不是channel的切片
	// 这里要给够缓冲长度capacity，否则会deadlock
	chs := make(chan int, len(list))
	// chs := make(chan int)
	fmt.Println("chs length =", len(chs))
	fmt.Println("chs capacity =", cap(chs))

	var wg sync.WaitGroup
	for i := 0; i < len(list); i++ {
		wg.Add(1)
		go func(inner InnerList, ch chan int, wg *sync.WaitGroup) {
			var ret int
			for _, x := range inner {
				ret += x
			}
			time.Sleep(1 * time.Second)
			ch <- ret
			wg.Done()
		}(list[i], chs, &wg)
	}

	wg.Wait()
	close(chs)

	for ret := range chs {
		fmt.Println("list result is:", ret)
	}
	fmt.Println("excute end.")
}

// generate a 3 * 10 two-dimensional slice
func genList() OuterList {
	var len int = 10
	inner1 := make([]int, len)
	inner2 := make([]int, len)
	inner3 := make([]int, len)

	// rand.Intn 是伪随机，需要先调用Seed
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		inner1[i] = rand.Intn(1000)
		inner2[i] = rand.Intn(1000)
		inner3[i] = rand.Intn(1000)
	}

	var outer OuterList = make([]InnerList, 0)
	outer = append(outer, inner1, inner2, inner3)
	return outer
}
