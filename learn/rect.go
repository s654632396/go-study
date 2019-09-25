package main

import (
	"fmt"
)

func main() {

	// r := &Rect{1, 2, 100, 200}
	r := NewRect(1, 2, 100, 150)
	fmt.Printf("the area is : %v", r.Area())
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func NewRect(x, y, width, height float64) *Rect {
	println("new rect")
	return &Rect{x, y, width, height}
}
