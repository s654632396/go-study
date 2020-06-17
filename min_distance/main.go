package main

import "fmt"

func main() {
	fmt.Println(
		minDistance("horse", "ros"),
	)
}

func minDistance(word1 string, word2 string) int {

	var dp [][]int
	var lw1, lw2 = len(word1), len(word2)
	dp = make([][]int, lw1)
	for i := 0; i < lw1; i++ {
		dp[i] = make([]int, lw2)
		dp[i][0] = i
	}
	for i := 0; i < lw2; i++ {
		dp[0][i] = i
	}
	

	return dp[lw1-1][lw2-1]
}

// 取最大值
func max(n ...int) int {
	var r int = n[0]
	for _, x := range n {
		if x > r {
			r = x
		}
	}
	return r
}
