package main

import (
	"fmt"
	"sync"
)

type Point struct {
	sync.Mutex
	x int
	y int
}

// call m1, p对象会进行一次内存拷贝
func (p Point) m1()  {
	fmt.Println(fmt.Sprintf("call m1, addr= %p", &p))
	p.x = 2 // will not change p.x
}

// call m2, receiver的p指针参数 即调用的对象参数的地址
func (p *Point) m2()  {
	fmt.Println(fmt.Sprintf("call m2, addr= %p", p))
	p.y = 2
}

func main()  {
	p1 := Point{x:1, y:1}
	fmt.Println(fmt.Sprintf("p1 addr= %p", &p1))
	p1.m1()
	p1.m2()
	fmt.Println("p=", p1)
}
