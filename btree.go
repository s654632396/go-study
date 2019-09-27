package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var items []int64 = []int64{29, 1, 35, 32, 11, 34, 54, 12, 2, 5, 9, 6, 45, 17, 33}
	var bt = NewBTree()
	for _, item := range items {
		bt.Add(item)
	}
	fmt.Printf("btree max node: %#v \n", bt.Max().value)
	fmt.Printf("btree min node: %#v \n", bt.Min().value)
	rand.Seed(int64(time.Now().Nanosecond()))
	var dist int64 = int64(rand.Intn(100))
	fmt.Printf("find %d in tree, %v\n", dist, bt.Find(dist))
	var f = func(n *Node) {
		fmt.Printf("%d\t", n.value)
	}
	var list []int64
	f = func(n *Node) {
		list = append(list, n.value)
	}
	bt.InOrderTraversal(f)
	fmt.Printf("\n in-order traverse binary tree -----------------------------\n%#v\n", list)
	list = nil
	bt.PreOrderTraversal(f)
	fmt.Printf("\n pre-order traverse binary tree-----------------------------\n%#v\n", list)
	list = nil
	bt.PostOrderTraversal(f)
	fmt.Printf("\n post-order traverse binary tree-----------------------------\n%#v\n", list)
	list = nil
}

// Node B-Tree Node Struct
type Node struct {
	value int64
	left  *Node
	right *Node
}

// BTree Struct
type BTree struct {
	RootNode *Node
}

// NewBTree Create Binary Tree
func NewBTree() (btree *BTree) {
	btree = &BTree{}
	return
}

// Add : add a new node to binary tree
func (bt *BTree) Add(value int64) {
	var node Node = Node{value: value}
	if bt.RootNode == nil {
		bt.RootNode = &node
		return
	}

	bt.RootNode.Insert(&node)
	// fmt.Println(value)
}

//Insert : insert node
func (node *Node) Insert(n *Node) {

	// fmt.Printf("node.value=%d,node.left=%v,right=%v \n", node.value, node.left, node.right)
	// fmt.Printf("currentNode.value=%d \n", n.value)

	if n.value < node.value {
		// 把n插入node的左边
		if node.left == nil {
			node.left = n
		} else {
			node.left.Insert(n)
		}
	} else {
		// 把n插入node的右边
		if node.right == nil {
			node.right = n
		} else {
			node.right.Insert(n)
		}
	}
	// os.Exit(0)
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
			// 当前节点比目标值大， 则查找当前节点到左树
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
//left ->  root -> right
func (bt *BTree) InOrderTraversal(cb func(n *Node)) {
	inOrderTraverse(bt.RootNode, cb)
	// fmt.Printf("\n")
}

func inOrderTraverse(node *Node, cb func(n *Node)) {
	if node != nil {
		inOrderTraverse(node.left, cb)
		// fmt.Printf("%d  ", node.value)
		cb(node)
		inOrderTraverse(node.right, cb)
	}
}

// PreOrderTraversal 先序遍历树
// root -> left -> right
func (bt *BTree) PreOrderTraversal(cb func(n *Node)) {
	preOrderTraverse(bt.RootNode, cb)
	// fmt.Printf("\n")
}

func preOrderTraverse(node *Node, cb func(n *Node)) {
	if node != nil {
		// fmt.Printf("%d  ", node.value)
		cb(node)
		preOrderTraverse(node.left, cb)
		preOrderTraverse(node.right, cb)
	}

}

// PostOrderTraversal 后续遍历树
// right -> left -> root
func (bt *BTree) PostOrderTraversal(cb func(n *Node)) {
	postOrderTraverse(bt.RootNode, cb)
	// fmt.Printf("\n")
}

func postOrderTraverse(node *Node, cb func(n *Node)) {
	if node != nil {
		postOrderTraverse(node.right, cb)
		postOrderTraverse(node.left, cb)
		// fmt.Printf("%d  ", node.value)
		cb(node)
	}
}
