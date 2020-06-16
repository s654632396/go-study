package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{3, 2, 1, 4, -1, 12}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
	fmt.Println(maxProduct([]int{0, 2}))
	fmt.Println(maxProduct([]int{3, -1, 4}))
}

func maxProduct(nums []int) int {
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		var res = 1
		for j := i; j < len(nums); j++ {
			res = res * nums[j]
			if max < res {
				max = res
			}
		}
	}
	return max
}