package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(123454321))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	max := int(math.Floor(math.Log10(float64(x)))) + 1

	var arr = make([]int, max/2)
	for i := 1; i <= max; i++ {
		if max%2 == 1 && (max/2+1 == i) {
			continue
		}
		n := x % int(math.Pow10(i))
		x -= n

		if i > max/2 {
			if n/int(math.Pow10(i-1)) != arr[max-i] {
				return false
			}
		} else {
			arr[i-1] = n / int(math.Pow10(i-1))
		}
	}
	return true
}
