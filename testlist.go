package main

import (
	"fmt"
	doublelist "github.com/s654632396/doublelist"
)

func main() {

	l := doublelist.NewDoubleList()

	l.RPush("aaaa")
	l.RPush("bbbbb")
	l.RPush("ccccc")
	l.RPush("poi")
	l.RPush("xxxx")
	l.RPush("shimakaze")

	// fmt.Printf("%v \n", l)

	fmt.Println(l.Index(3).Value())
	fmt.Println(l.Index(-3).Value())
	fmt.Println(l.Index(-1).Value())
	var s []*doublelist.ListNode
	var err error
	if s, err = l.Range(-5, 10); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println("range s[", i, "]", s[i].Value())
	}
	if s, err = l.Range(-5, -10); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(s); i++ {
		fmt.Println("range s[", i, "]", s[i].Value())
	}
}
