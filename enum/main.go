package main

import (
	"fmt"
)

type Pai int

func (p Pai) String() string {
	switch p {
	case 0:
		return "万"
	case 1:
		return "饼"
	case 2:
		return "索"
	case 3:
		return "字"
	default:
		panic(`invalid Pai`)
	}
}

const (
	Man Pai = iota
	Pin
	Sou
	Zii
)

func main() {

	fmt.Println(
		Man,
		Pin,
		Sou,
		Zii,
	)
	fmt.Println(
		Man == 0,
		Pin == 1,
		Sou == 2,
		Zii == 3,
	)
}
