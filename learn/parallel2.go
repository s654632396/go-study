package main

import (
	"fmt"
	"math/rand"
	"runtime"
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
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			fmt.Printf("%d ", list[i][j])
		}
		fmt.Println(" ")
	}
	fmt.Println("-----------------------------------------")

	fmt.Println(len(list), len(list[0]), len(list[1]), len(list[2]))
	// ===== 使用带缓冲区的channel，而不是channel的切片
	chs := make(chan int, len(list))
	fmt.Println("chs length =", len(chs))
	fmt.Println("chs capacity =", cap(chs))

	for i := 0; i < len(list); i++ {
		i := i
		go func(inner InnerList, ch chan int) {
			var ret int
			for _, x := range inner {
				ret += x
			}
			ch <- ret
		}(list[i], chs)
	}

	for i := 0; i < cap(chs); i++ {
		ret := <-chs
		fmt.Println("list result is:", ret)
	}

}

// generate a 3 * 10 two-dimensional slice
func genList() OuterList {
	var len int = 10
	inner1 := make([]int, len)
	inner2 := make([]int, len)
	inner3 := make([]int, len)
	for i := 0; i < len; i++ {
		inner1[i] = rand.Intn(1000)
		inner2[i] = rand.Intn(1000)
		inner3[i] = rand.Intn(1000)
	}

	var outer OuterList = make([]InnerList, 0)
	outer = append(outer, inner1, inner2, inner3)
	return outer
}
