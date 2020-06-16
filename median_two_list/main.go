package main

import (
	"fmt"
	"math"
)

/**
 * 寻找两个正序数组的中位数
 * 	@see https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
 */

func main() {

	n1 := []int{1, 4}
	n2 := []int{2, 3}

	r := findMedianSortedArrays(n1, n2)

	fmt.Println(r)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := make([]int, 0)
	// merge two sorted list to one sorted list
	l1,l2 := len(nums1), len(nums2)
	l := l1 + l2
	for i, j, k := 0, 0, 0; i < l; i++ {
		if j == len(nums1) {
			num = append(num, nums2[k])
			k++
			continue
		}
		if k == len(nums2) {
			num = append(num, nums1[j])
			j++
			continue
		}

		if nums1[j] >= nums2[k] {
			num = append(num, nums2[k])
			k++
		} else {
			num = append(num, nums1[j])
			j++
		}
	}
	var mid float64
	if l%2 == 0 {
		mid = (float64(num[l/2]) + float64(num[l/2-1])) / 2
	} else {
		mid = float64(num[int(math.Floor(float64(l/2)))])
	}

	return mid
}
