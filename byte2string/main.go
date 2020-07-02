package main

import (
	"log"
	"unsafe"
)

func main() {
	bs := make([]byte, 100)
	bs = []byte{'a', 'b', 'c', 'd', 'e'}
	s1 := string(bs)
	log.Println(s1)
	// 显示转换指针类型,然后取指针变量的值赋值给s2
	s2 := *(*string)(unsafe.Pointer(&bs))
	//log.Println(s2)
	//bs[0] = 'Q' // @!! 修改bs的元素会影响到s2
	//log.Println(s2, string(bs))
	//bs = []byte{'Q', 'Q', 'Q'}
	//log.Println(s2, string(bs))
	log.Printf("&s2=%p, bs=%p\n", &s2, bs)
}
