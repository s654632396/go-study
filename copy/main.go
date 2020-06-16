package main

import (
	"fmt"
)

func main() {

	// correct
	var a1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 = make([]int, 5)
	copy(s1, a1[:5])
	var s2 = make([]int, 5)
	copy(s2, a1[4:])
	fmt.Println(s1, s2, a1)
	copy(s1[2:], s2)
	fmt.Println(s1, s2, a1)
	fmt.Println(makeSlice(1, 2, 3, 47, 8, 9, 0))

	// wrong
	var s11 = makeSlice(10, 20, 30, 40)
	var s12 = s11[1:3]
	fmt.Println(s11, s12)
	s12[1] = 90
	fmt.Println(s11, s12)
	var s13 = s12[0:1]
	fmt.Println(s11, s12, s13)
	s13[0] = 100
	fmt.Println(s11, s12, s13)

	// correct
	var s14 = makeSlice(s11[1:3]...)
	fmt.Println(s11, s14)
	s14[1] = 0
	fmt.Println(s11, s14)

}

// safely create slice by copy
func makeSlice(data ...int) []int {
	var list = make([]int, len(data))
	copy(list, data)
	return list
}
