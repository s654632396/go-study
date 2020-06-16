package main

import (
	"fmt"
	"log"
	"math"
)

/**
 * https://leetcode-cn.com/problems/climbing-stairs/solution/
 */

func main() {

	// fmt.Println(climbStairsFib(5))

	for i := 1; i < 47; i++ {
		r1 := climbStairsFib(i)
		r2 := climbStairs(i)
		fmt.Println(i, "-----", r1, r2)
		if r1 != r2 {
			log.Printf("err: i=%d, r1=%d, r2=%d\n", i, r1, r2)
		}
	}
}

// ********* 递归 + cache **************
var cacheMap = make(map[int]int)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if res, ok := cacheMap[n]; ok {
		return res
	}
	r := climbStairs(n-1) + climbStairs(n-2)
	cacheMap[n] = r
	return r
}

// ********** 斐波那契数列通项公式 ***********

func climbStairsFib(n int) int {
	n = n + 1
	var r int

	var s5 float64 = math.Sqrt(5)

	// 1/s5[((1+s5)/2)^n - ((1-s5)/2)^n]
	var p1 = math.Pow((1+s5)/2, float64(n))
	var p2 = math.Pow((1-s5)/2, float64(n))
	r = int(math.Round((p1 - p2) / s5))

	return r
}
