package main

/*
 * leetcode: https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/submissions/
 * 思路: 顺时针遍历, 边界换方向
 * 此题还有另外一种思路: 矩阵逆时针旋转,取第一行
 */
import (
	"fmt"
	"log"
)

func main() {

	matrix := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16}}

	fmt.Println(spiralOrder(matrix))
}

type mPtr struct {
	x int // rowIndex
	y int // colIndex
}

type Director struct {
	v      int // value range of 0~3
	w      int // width
	h      int // height
	p      *mPtr
	border []int // {0, 0, mw, mh}
}

func NewDirector(w, h int) (d *Director) {
	d = new(Director)
	d.w, d.h = w, h
	d.border = []int{0, 0, w - 1, h - 1}
	return
}

func (d *Director) next() (y, x int) {
	if d.p == nil {
		// start point is 0,0
		d.p = new(mPtr)
		d.p.x, d.p.y = 0, 0
		d.v = 0
	} else {
		// log.Printf("d.v=%d, x, y= %d, %d \n", d.v, d.p.x, d.p.y)
	BEGIN:
		switch d.v {
		case 0:
			// move to right, incr p.x if p.x less than rowIndex - 1
			if d.p.x < d.w-1 && !d.borderHit(d.p.x, d.p.y) {
				d.p.x++
			} else {
				d.v++
				d.border[1]++
				goto BEGIN
			}
		case 1:
			// move to bottom, incr p.y if p.y less than colIndex - 1
			if d.p.y < d.h-1 && !d.borderHit(d.p.x, d.p.y) {
				d.p.y++
			} else {
				d.v++
				d.border[2]--
				goto BEGIN
			}
		case 2:
			// move to right, decr p.x if p.x larger than 0
			if d.p.x > 0 && !d.borderHit(d.p.x, d.p.y) {
				d.p.x--
			} else {
				d.v++
				d.border[3]--
				goto BEGIN
			}
		case 3:
			// move to top, decr p.y if p.y larger than 0
			if d.p.y > 0 && !d.borderHit(d.p.x, d.p.y) {
				d.p.y--
			} else {
				d.v = 0 // reset
				d.border[0]++
				goto BEGIN
			}
		}
	}
	return d.p.x, d.p.y
}

func (d *Director) borderHit(x, y int) (r bool) {
	switch d.v {
	case 0: //move to right, border is right, border[2]
		if x == d.border[2] {
			r = true
		}
	case 1: //move to bottom, border is bottom, border[3]
		if y == d.border[3] {
			r = true
		}
	case 2: //move to left, border is bottom, border[0]
		if x == d.border[0] {
			r = true
		}
	case 3: //move to up, border is bottom, border[1]
		if y == d.border[1] {
			r = true
		}
	}
	return
}

func spiralOrder(matrix [][]int) []int {
	log.Println(matrix)
	// 测量长宽
	colLen := len(matrix)

	// order init
	order := make([]int, 0)
	if colLen == 0 {
		return order
	}
	rowLen := len(matrix[0])

	// check each row
	for i := 1; i < colLen; i++ {
		if len(matrix[i]) != rowLen {
			panic("invalid input matrix!")
		}
	}

	// new director
	d := NewDirector(rowLen, colLen)

	for s := 0; s < rowLen*colLen; s++ {
		y, x := d.next()
		order = append(order, matrix[x][y])
	}

	return order
}
