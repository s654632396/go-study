package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {

	var trangle = [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	fmt.Println(minimumTotal(trangle))

}

// https://leetcode-cn.com/problems/triangle/
// 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	var sum int
	var depth = len(triangle)
	var dp = make([][]int, depth)

	return sum
}

// 全排列
func permute(nums []int) [][]int {
	var all = make([][]int, 0)
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	var cpNums = make([]int, len(nums))
	copy(cpNums, nums)

	var idx = 0
	for i := 0; i < len(nums); i++ {
		cpNums[0], cpNums[i] = cpNums[i], cpNums[0]
		subAll := permute(cpNums[1:])
		for j := 0; j < len(subAll); j++ {
			tmp := make([]int, len(nums))
			tmp[0] = cpNums[0]
			copy(tmp[1:], subAll[j])
			all = append(all, make([]int, len(nums)))
			copy(all[idx], tmp)
			idx++
		}
	}

	return all
}

func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		isF, isB := i%3 == 0, i%5 == 0
		if !isF && !isB {
			ret[i-1] = strconv.Itoa(i)
		} else if isF && !isB {
			ret[i-1] = "Fizz"
		} else if !isF && isB {
			ret[i-1] = "Buzz"
		} else {
			ret[i-1] = "FizzBuzz"
		}
	}
	return ret
}

// 埃拉托斯特尼筛法
// 参考: https://zh.wikipedia.org/wiki/%E5%9F%83%E6%8B%89%E6%89%98%E6%96%AF%E7%89%B9%E5%B0%BC%E7%AD%9B%E6%B3%95
func countPrimes(n int) int {
	var c int
	if n <= 2 {
		return c
	}
	if n == 3 {
		return 1
	}
	var arr = make([]bool, n+1)
	arr[0], arr[1] = true, true

	for x := 2; x <= int(math.Sqrt(float64(n))); x++ {
		if !arr[x] {
			c++
			for i := x * x; i < n; i += x {
				arr[i] = true
			}
		}
	}
	for x := 1 + int(math.Sqrt(float64(n))); x < n; x++ {
		if !arr[x] {
			c++
		}
	}
	return c
}

// exceed time!!
func countPrimesV1(n int) int {
	var c int
E:
	for x := 2; x < n; x++ {
		for i := 2; i < x; i++ {
			if x%i == 0 {
				continue E
			}
		}
		log.Println(x)
		c++
	}
	return c
}
