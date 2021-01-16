package main

import "fmt"

type Money struct {
	amount float64
}

func main() {
	//var a = Money{}
	//var b = &Money{}
	//var c = new(Money)
	//var d Money
	//var e *Money
	//fmt.Println(a, b, c, d, e)

	var s1 = []int{1,2,3,4,5}
	var s2 = []int{6,7,8,9}


	var s3 = append(s1, s2...)
	var s4 = make([]int, len(s1) + len(s2))
	copy(s4[0:], s1)
	copy(s4[len(s1):], s2)
	fmt.Println(s3, s4)

}
