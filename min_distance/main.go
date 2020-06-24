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
	dp = make([][]int, lw1+1)
	for i := 0; i <= lw1; i++ {
		dp[i] = make([]int, lw2+1)
		dp[i][0] = 1
	}
	for i := 0; i <= lw2; i++ {
		dp[0][i] = 1
	}
	for i := 0; i <= lw1-1; i++ {
		for j := 0; j <= lw2-1; j++ {
			fmt.Println( string(word1[i]), string(word2[j]), dp[i+1][j], dp[i][j+1])
			if word1[i] != word2[j] {
				dp[i+1][j+1] = dp[i][j+1]+dp[i+1][j]
			}
		}
	}

	return dp[lw1][lw2]
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
