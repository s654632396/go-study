package main

import (
	"fmt"
	"time"
)

func main() {

	for c := 10; c > 0; c-- {
		var t1 = time.Now()
		test_loop_a()
		fmt.Println(
			"loop a cost time:" + time.Now().Sub(t1).String(),
		)
		var t2 = time.Now()
		test_loop_b()
		fmt.Println(
			"loop b cost time:" + time.Now().Sub(t2).String(),
		)
	}

}

func test_loop_a() {
	var arr = make([]int, 64*1024*1024)
	for i := 0; i < len(arr); i++ {
		arr[i] *= 3
	}
}

func test_loop_b() {
	var arr = make([]int, 64*1024*1024)
	for i := 0; i < len(arr); i += 16 {
		arr[i] *= 3
	}
}
