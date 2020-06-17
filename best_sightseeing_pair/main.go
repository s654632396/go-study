package main

import (
	. "fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)
/**
 * 最佳观光组合
 * https://leetcode-cn.com/problems/best-sightseeing-pair/submissions/
 */

func main() {
	f, err := os.Open("/tmp/2")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	content, _ := ioutil.ReadAll(f)
	c := string(content)
	c = strings.Trim(c, "[] ")
	cl := strings.Split(c, ", ")
	var data = make([]int, len(cl))
	for i, s := range cl {
		data[i], _ = strconv.Atoi(strings.Trim(s, " \t"))
	}
	Println(
		maxScoreSightseeingPair(data),
		//maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}),
	)

}

func maxScoreSightseeingPair(A []int) int {

	l := len(A)
	var score int
	var maxI int = A[0] + 0

	for j := 1; j < l; j++ {
		score = max(maxI+A[j]-j, score)
		maxI = max(maxI, A[j]+j)
	}

	return score
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
