package main

import "fmt"

func main() {

	var list = []int{3, 3}
	var target int = 6

	var a = twoSum(list, target)
	fmt.Printf("a=%d, b=%d\n", a[0], a[1])
}

func twoSum(list []int, target int) []int {
	var m = make(map[int][]int)
	var result []int = make([]int, 2)
	for idx, item := range list {
		if _, ok := m[target-item]; !ok {
			m[target-item] = []int{item, idx}
		}
		if x, ok := m[item]; ok {
			if x[0]+item == target && x[1] != idx {
				result[0], result[1] = x[1], idx
			}
		}
	}

	return result
}
