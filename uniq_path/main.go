package main

import "fmt"

/**
 * 不同路径
 * https://leetcode-cn.com/problems/unique-paths/submissions/
 */

func main() {
	fmt.Println(uniquePathsDp(7, 3))
}

// ***** 递归,自顶向下, 将大问题拆分小问题, 小问题重复求解用map缓存 *****

var cacheDict = make(map[string]int)

func uniquePaths(m int, n int) int {

	if m == 1 || n == 1 {
		return 1
	}
	if m > 100 || n > 100 {
		panic(`invalid input`)
	}
	var k = string(m) + "-" + string(n)
	var res int
	if s, ok := cacheDict[k]; ok {
		res = s
	} else {
		res = uniquePaths(m-1, n) + uniquePaths(m, n-1)
		cacheDict[k] = res
	}

	return res
}


// ***** 动态规划,自下向上 *****

func uniquePathsDp(m int, n int) int {
	var arr [][]int
	arr = make([][]int, n)
	for i := 1; i <= n; i++ {
		if len(arr[i-1]) == 0 {
			arr[i-1] = make([]int, m)
		}

		arr[i-1][0] = 1
	}
	for i := 1; i <= m; i++ {
		arr[0][i-1] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			arr[i][j] = arr[i-1][j] + arr[i][j-1]
		}
		fmt.Println(arr[i])
	}

	return arr[n-1][m-1]
}
