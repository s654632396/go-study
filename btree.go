package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var bt = NewBTree()
	// var items []int64 = []int64{29, 1, 35, 32, 11, 34, 54, 12, 2, 5, 9, 6, 45, 17, 33}
	// for _, item := range items {
	// 	bt.Add(item)
	// }
	var seed int = time.Now().Nanosecond()
	rand.Seed(int64(seed))
	for i := 0; i < 80; i++ {
		value := int64(rand.Intn(100))
		bt.Add(value, value+100000)
	}
	// fmt.Printf("btree max node: %#v \n", bt.Max().value)
	// fmt.Printf("btree min node: %#v \n", bt.Min().value)
	// var dist int64 = int64(rand.Intn(100))
	// src := bt.Find(dist)
	// fmt.Printf("find %d in tree,value=%v,item=%v \n", dist, src.value, src.item)
	// fmt.Printf("BinaryTree len: %d \n", bt.Len)
	bt.String()

	var list []int64
	var f1 = func(n *Node) {
		list = append(list, n.value)
	}
	bt.InOrderTraversal(f1)
	fmt.Printf("in-order traverse binary tree -----------------------------\n%#v\n", list)
	return
	list = nil
	bt.PreOrderTraversal(f1)
	fmt.Printf("\n pre-order traverse binary tree-----------------------------\n%#v\n", list)
	list = nil
	bt.PostOrderTraversal(f1)
	fmt.Printf("\n post-order traverse binary tree-----------------------------\n%#v\n", list)
	list = nil
}

// Item : node storage values
type Item interface{}

// Node B-Tree Node Struct
type Node struct {
	value int64
	item  Item
	left  *Node
	right *Node
	pos   int8 // 0 tree-root 1 left 2 right
}

// BTree Struct
type BTree struct {
	RootNode *Node
	Len      int
}

// NewBTree Create Binary Tree
func NewBTree() (btree *BTree) {
	btree = &BTree{}
	return
}

// Add : add a new node to binary tree
func (bt *BTree) Add(value int64, item Item) {
	var node Node = Node{value: value, item: item}
	if bt.RootNode == nil {
		bt.RootNode = &node
		return
	}

	if bt.RootNode.Insert(&node) {
		bt.Len++
	}
}

// Insert : insert node
func (node *Node) Insert(n *Node) (b bool) {
	b = false
	if n.value < node.value {
		//  把n插入node的左边
		if node.left == nil {
			n.pos = 1 // mark node as left Node
			node.left = n
		} else {
			node.left.Insert(n)
		}
		b = true
		return
	}
	if n.value > node.value {
		//  把n插入node的右边
		if node.right == nil {
			n.pos = 2 // mark node as right Node
			node.right = n
		} else {
			node.right.Insert(n)
		}
		b = true
		return
	}
	b = false
	return
}

// Min : return minium node of binary tree
func (bt *BTree) Min() (node *Node) {
	node = bt.RootNode
	for {
		if node.left != nil {
			node = node.left
			continue
		}
		return node
	}
}

// Max : return Maxium node of binary tree
func (bt *BTree) Max() (node *Node) {
	node = bt.RootNode
	for {
		if node.right != nil {
			node = node.right
			continue
		}
		return node
	}
}

// Find : find Node By value
func (bt *BTree) Find(value int64) (node *Node) {
	node = bt.RootNode
	for {
		if node.value == value {
			return
		}
		if node.value > value {
			//  当前节点比目标值大， 则查找当前节点到左树
			node = node.left
		} else {
			node = node.right
		}
		if node == nil {
			return
		}
	}
}

// InOrderTraversal 中序遍历树
// left ->  root -> right
func (bt *BTree) InOrderTraversal(cb func(n *Node)) {
	inOrderTraverse(bt.RootNode, cb)
}

func inOrderTraverse(node *Node, cb func(n *Node)) {
	if node != nil {
		inOrderTraverse(node.left, cb)
		cb(node)
		inOrderTraverse(node.right, cb)
	}
}

// PreOrderTraversal 先序遍历树
//  root -> left -> right
func (bt *BTree) PreOrderTraversal(cb func(n *Node)) {
	preOrderTraverse(bt.RootNode, cb)
}

func preOrderTraverse(node *Node, cb func(n *Node)) {
	if node != nil {
		cb(node)
		preOrderTraverse(node.left, cb)
		preOrderTraverse(node.right, cb)
	}

}

// PostOrderTraversal 后续遍历树
//  right -> left -> root
func (bt *BTree) PostOrderTraversal(cb func(n *Node)) {
	postOrderTraverse(bt.RootNode, cb)
}

func postOrderTraverse(node *Node, cb func(n *Node)) {
	if node != nil {
		postOrderTraverse(node.right, cb)
		postOrderTraverse(node.left, cb)
		cb(node)
	}
}

// LevelOrderTraversal : 层次遍历
func (bt *BTree) LevelOrderTraversal(cb func(n *Node, lv int, idx int, t int)) {
	node := bt.RootNode
	var nodes []*Node = make([]*Node, 0)
	nodes = append(nodes, node)
	levelOrderTraverse(nodes, 0, cb)
}

func levelOrderTraverse(nodes []*Node, level int, cb func(n *Node, lv int, idx int, t int)) {
	var nextNodes []*Node = make([]*Node, 0)
	var nodeLen = len(nodes)
	for index, node := range nodes {
		if node.left != nil {
			nextNodes = append(nextNodes, node.left)
		}
		if node.right != nil {
			nextNodes = append(nextNodes, node.right)
		}
		cb(node, level, index, nodeLen)
	}
	level++
	if len(nextNodes) > 0 {
		levelOrderTraverse(nextNodes, level, cb)
	}
}

// NodeList :
type NodeList []*Node

// String 输出二叉树
func (bt *BTree) String() {
	var debug []string
	defer func() {
		// debug output binary tree
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
			fmt.Printf("\n%+v\n", debug)
		}
	}()
	var NLists = make([]NodeList, 0)
	var list []*Node
	var depth int
	var f0 = func(n *Node, lv int, idx int, lt int) {
		list = append(list, n)
		if idx+1 == lt {
			NLists = append(NLists, list)
			list = nil
		}
		depth = lv
	}
	bt.LevelOrderTraversal(f0)
	fmt.Printf(">>> BINARY TREE___[total node:%d, depth:%d]______________________________________________________________\n", bt.Len, depth)
	var blank = " "
	strings.Repeat(blank, 1)
	var lastTab int
	var lineChars int
	for _, nodes := range NLists {
		fmt.Println()
		for _, node := range nodes {
			lineChars = len(strconv.Itoa(int(node.value)))
			tc := int(node.value) - lastTab - lineChars + 1
			lastTab = tc + lastTab
			debug = append(debug, fmt.Sprintf("node=%d,tc=%d,lastTab=%d\n", node.value, tc, lastTab))
			tab := strings.Repeat(blank, tc)
			if node.pos == 1 {
				fmt.Printf("%s<%d", tab, node.value)
			} else if node.pos == 2 {
				fmt.Printf("%s%d>", tab, node.value)
			} else {
				// root
				fmt.Printf("%s<%d>", tab, node.value)
			}
		}
		fmt.Println()
		lastTab = 0
	}
	fmt.Printf("\n <<< BINARY TREE_________________________________________________________________\n")
}
