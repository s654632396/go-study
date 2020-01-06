package main

import (
	"fmt"
	"math/rand"
	"time"

	adt "github.com/s654632396/mpkg/btree"
)

func main() {
	var bt = adt.NewBTree()
	var items []adt.NodeValue = []adt.NodeValue{29, 1, 35, 32, 11, 34, 54, 12, 2, 5, 9, 6, 45, 17, 33}
	for _, item := range items {
		bt.Add(adt.NodeValue(item), item+10000)
	}
	var seed int = time.Now().Nanosecond()
	rand.Seed(int64(seed))
	fmt.Println("start ..")
	// for i := 0; i < 80; i++ {
	// 	value := NodeValue(rand.Intn(100))
	// 	bt.Add(value, value+100000)
	// }

	var list []adt.NodeValue
	var f1 = func(n *adt.Node) {
		list = append(list, n.Value())
	}
	bt.InOrderTraversal(f1)
	fmt.Printf("in-order traverse binary tree -----------------------------\n%+v\n", list)

	bt.String()
	unlink := bt.Find(29)
	if unlink != nil {
		unlinked := bt.Unlink(unlink)
		fmt.Printf("SUCCESSFULLY UNLINKED NODE: %d\n", unlinked)
		bt.String()
		bt.InOrderTraversal(f1)
		fmt.Printf("in-order traverse binary tree -----------------------------\n%+v\n", list)
	}

	return
}
