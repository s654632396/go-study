package main

import (
	"fmt"
)

/**
 * 最长连续序列
 */
func main() {
	//r1 := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	//r2 := longestConsecutive([]int{1, 2, 0, 1})
	//r3 := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	r3 := longestConsecutive2([]int{-1,9,-3,-6,7,-8,-6,2,9,2,3,-2,4,-1,0,6,1,-9,6,8,6,5,2})
	/**/
	/**/
	//r4 := longestConsecutive([]int{0, -1})

	fmt.Println(r3)

}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var max = 1                 // 返回值,最大连续长度
	var m = make(map[int][]int) // map[number]区间[begin, end]
	for i := 0; i < len(nums); i++ {
		if _, hitS := m[nums[i]]; hitS {
			continue // 重复数字,跳过
		}
		var number = nums[i]
		m[number] = []int{number, number} // 初始化当前数字
		_, hitL := m[number-1]
		_, hitR := m[number+1]
		if hitL && hitR {
			left, right := m[number-1][0], m[number+1][len(m[number+1])-1]
			m[number] = []int{left, right}
			m[left], m[right] = m[number], m[number]
		} else if hitL {
			left, right := m[number-1][0], number
			m[number] = []int{left, right}
			m[left], m[right] = m[number], m[number]
		} else if hitR {
			left, right := number, m[number+1][len(m[number+1])-1]
			m[number] = []int{left, right}
			m[left], m[right] = m[number], m[number]
		}
		if m[number][len(m[number])-1]-m[number][0]+1 > max {
			max = m[number][len(m[number])-1] - m[number][0] + 1
		}
	}
	return max
}


func longestConsecutive2(nums []int) int {
	d, ans := make(map[int]bool, len(nums)), 0
	for _, i := range nums {
		d[i] = true
	}
	for _, i := range nums {
		if !d[i - 1] {
			t := 1
			for i++; d[i]; i++ {t++}
			if t > ans {
				ans = t
			}
		}
	}
	return ans
}
