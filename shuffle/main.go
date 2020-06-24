package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor([]int{1, 2, 3})
	param_1 := obj.Reset()
	for i := 0; i < 10; i++ {
		param_2 := obj.Shuffle()
		fmt.Println(param_1, param_2)
	}
}

type Solution struct {
	origins  []int
	currents []int
	l        int
}

func Constructor(nums []int) Solution {
	s := Solution{nums, nil, len(nums)}
	s.currents = make([]int, s.l)
	copy(s.currents, s.origins)
	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origins
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	// generate rand 0~l length []int keys
	for i := 0; i < this.l; i++ {
		// generate a key to be exchange with i
		k := rand.Intn(this.l - i) + i
		if i != k {
			this.currents[i], this.currents[k] = this.currents[k], this.currents[i]
		}
	}
	return this.currents
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
