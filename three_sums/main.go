package main

import (
	"fmt"
	"sort"
)

/**
 * 三数之和
 * https://leetcode-cn.com/problems/3sum/submissions/
 * 解题思路: 排序 + 双指针
 */

func main() {
	fmt.Println(
		threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}),
	)
}

func threeSum(nums []int) [][]int {
	var ret [][]int
	ret = make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var sum int
		for j, k := i+1, len(nums)-1; j < k; {
			sum = nums[i] + nums[j] + nums[k] // Q: 如何处理大整数之和溢出的情况?
			if sum > 0 {
				k--
				continue
			}
			if sum < 0 {
				j++
				continue
			}
			if len(ret) > 0 && ret[len(ret)-1][0] == nums[i] && ret[len(ret)-1][1] == nums[j] {
				j++
				continue
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
			}
		}
	}
	return ret
}
