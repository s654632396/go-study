package main

import (
	"container/list"
	"fmt"
	"log"
)

func main() {

	l := list.New()
	l.PushFront("a")
	lb := l.PushBack("b")
	le := l.PushBack("e")
	l.PushBack("f")

	l.InsertAfter("c", lb)
	l.InsertBefore("d", le)

	for node := l.Front(); node != nil; node = node.Next() {
		fmt.Println(node.Value)
	}

	log.Println(l.Len())

	l.Init()
	log.Println(l.Len())
}
