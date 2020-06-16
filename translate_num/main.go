package main

import (
	"fmt"
	"strconv"
)

/**
 * https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
 */

func main() {

	translateNum(12258)
	//translateNum(25)

}


func translateNum(num int) int {

	ns := strconv.Itoa(num)

	fmt.Println(ns)

	return 0
}

