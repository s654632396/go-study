package main

import (
	"fmt"
)

func Strstr(haystack, needle string) (ret int) {
	needleLen := len(needle)

	ret = -1

	for i, _ := range haystack {

		cmpstr := haystack[i : needleLen+i]
		if cmpstr == needle {
			ret = i
			break
		}
	}

	return
}

func main() {

	haystack := "hello"
	needle := "ll"

	fmt.Println(Strstr(haystack, needle))
}
