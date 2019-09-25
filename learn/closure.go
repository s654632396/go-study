package main

import (
	"fmt"
)

func main() {
	func(a int) {
		fmt.Println("a = ", a)
	}(123)
}
