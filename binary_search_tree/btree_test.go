package btree

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestBTree_Main(t *testing.T) {
	startTime := time.Now()
	var bt = NewBTree()
	var items []NodeIndex = []NodeIndex{29, 1, 35, 32, 11, 34, 54, 12, 2, 5, 9, 6, 45, 17, 33}
	for _, item := range items {
		bt.Add(item, item+10000)
	}
	var seed = time.Now().Nanosecond()
	rand.Seed(int64(seed))
	for i := 0; i < 80; i++ {
		value := NodeIndex(rand.Intn(100))
		bt.Add(value, value+100000)
	}
	log.Println(fmt.Sprintf("insert 100 node use time: %v", time.Now().Sub(startTime)))
	startTime = time.Now()
	var list1 []NodeIndex
	var f1 = func(n *Node) {
		list1 = append(list1, n.Index())
	}
	bt.InOrderTraversal(f1)
	log.Println(fmt.Sprintf("in order traveling all 100 node use time: %v", time.Now().Sub(startTime)))
	log.Printf("in-order traverse binary tree -----------------------------\n%+v\n", list1)

	// unlink a node
	startTime = time.Now()
	for i := 50; i < 60; i++ {
		unlink := bt.Find(NodeIndex(i)) // search node to be unlinked by index
		if unlink != nil {
			// log.Println(unlink.Item())
			_ = bt.Unlink(unlink)
		} else {
			log.Println("no hit unlink index node:", i)
		}
	}
	log.Println(fmt.Sprintf("unlink 10 node use time: %v", time.Now().Sub(startTime)))

	// view traveling nodes after unlink
	var list2 []NodeIndex
	var f2 = func(n *Node) {
		list2 = append(list2, n.Index())
	}
	bt.InOrderTraversal(f2)
	log.Printf("in-order traverse binary tree -----------------------------\n%+v\n", list2)

	return
}
