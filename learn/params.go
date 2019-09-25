package main

import (
	"fmt"
)

func main() {

	//	MyFunc("a", "b", "c")
	args := make([]string, 3)
	args = []string{"aa", "b", "c"}
	MyFunc(args...)
}

func MyFunc(p ...string) {
	for _, arg := range p {
		fmt.Println("arg :", arg)
	}

}
