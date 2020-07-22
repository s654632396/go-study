package main

import (
	"log"
	"unsafe"
)

type Sample struct {
	// environment: x86_64
	// size + padding = total
	A int8  // 1 + 3 = 4
	B int32 // 4 + 0 = 4
	X int16 // 2 + 6 = 8
	Y int64 // 8 + 0 = 8
	Z int64 // 8 + 0 = 8
	//  总长度	sum  = 32
}

// [summary] 对齐填充padding的计算:
// 如果后一个成员的对齐系数 <= 当前成员的长度
// 当前成员的padding=0
// 例如:
// B的padding(0) = X的对齐系数(2) < B的长度(4)
// Z的padding(0) = Y的对齐系数(8) = Z的长度(8)
//
// 如果后一个成员的对齐系数 > 当前成员的长度
// 当前成员的padding = 后一个成员的对齐系数 - 当前成员的长度
// 例如:
// A的padding(3) = B的对齐系数(4) - A的长度(1)

// 参考:
// https://ms2008.github.io/2019/08/01/golang-memory-alignment/

func main() {

	var str string
	log.Println(unsafe.Alignof(str)) // string的内存对齐长度为8byte

	s := new(Sample)
	// 获取s.X的指针
	pX := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(unsafe.Sizeof(s.A)) + 3 + unsafe.Sizeof(s.B)))
	*pX = int16(114)
	log.Println(s.X)

	log.Println(unsafe.Alignof(s.A), unsafe.Alignof(s.B), unsafe.Alignof(s.X), unsafe.Alignof(s.Y), unsafe.Alignof(s.Z))
	log.Println(unsafe.Offsetof(s.A), unsafe.Offsetof(s.B), unsafe.Offsetof(s.X), unsafe.Offsetof(s.Y), unsafe.Offsetof(s.Z))
	pY := (*int64)(
		unsafe.Pointer(
			// s.X的内存地址 + s.X的值内存长度 +  s.X的内存对齐长度 = 成员Y的内存地址
			uintptr(unsafe.Pointer(pX)) + unsafe.Sizeof(int16(0)) + 6,
		),
	)
	*pY = int64(514)
	log.Println(s.Y)

	pZ := (*int64)(
		unsafe.Pointer(
			uintptr(unsafe.Pointer(pY)) + unsafe.Sizeof(int64(0)),
		),
	)
	*pZ = int64(401)
	log.Println(s.Z)
	log.Println(unsafe.Sizeof(*s)) // 1+(3)+4+(0)+2+(6)+8+(0)+8+(0)=32

}
