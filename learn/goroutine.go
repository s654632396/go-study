package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}

func Add(x, y int) int {
	var ret int
	ret = x + y
	fmt.Println("ret = ", ret)
	return ret
}
