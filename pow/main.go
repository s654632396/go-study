package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(-1.00000, -2147483647))

}

func myPow(x float64, n int) float64 {
	var isNegative bool
	var isOdd bool
	if n < 0 {
		isNegative = true
		n = -n
	}
	if n%2 != 0 {
		isOdd = true
	}
	if n == 0 || x == 1 {
		return 1
	}
	if x == -1 {
		if isOdd {
			return -1
		} else {
			return 1
		}
	}
	var ret float64
	if isOdd {
		ret = x
	} else {
		ret = 1
	}
	x2 := x * x
	for i := 0; i < n/2; i++ {
		if ret == 0 {
			return 0
		}
		if math.IsInf(ret, 0) {
			ret = math.Inf(0)
			break
		}
		ret *= x2
	}
	if isNegative {
		return 1 / ret
	}
	return ret
}
