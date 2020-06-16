package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(
		fourSum([]int{-1,0,1,2,-1,-4}, -1),
	)
}

func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	ret = make([][]int, 0)
	if len(nums) < 4 {
		return ret
	}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for p := i + 1; p < len(nums); p++ {
			if p > i+1 && nums[p] == nums[p-1] {
				continue
			}
			var sum int
			for j, k := p+1, len(nums)-1; j < k; {
				sum = nums[i] + nums[p] + nums[j] + nums[k] // Q: 如何处理大整数之和溢出的情况?
				if sum > target {
					k--
					continue
				}
				if sum < target {
					j++
					continue
				}
				if len(ret) > 0 && ret[len(ret)-1][0] == nums[i] && ret[len(ret)-1][1] == nums[p] && ret[len(ret)-1][2] == nums[j] {
					j++
					continue
				} else {
					ret = append(ret, []int{nums[i], nums[p], nums[j], nums[k]})
					j++
				}
			}
		}

	}

	return ret
}
