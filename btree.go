package main

import (
	"fmt"
	"math/rand"
	// "os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var bt = NewBTree()
	var items []NodeValue = []NodeValue{29, 1, 35, 32, 11, 34, 54, 12, 2, 5, 9, 6, 45, 17, 33}
	for _, item := range items {
		bt.Add(NodeValue(item), item+10000)
	}
	var seed int = time.Now().Nanosecond()
	rand.Seed(int64(seed))
	// for i := 0; i < 80; i++ {
	// 	value := NodeValue(rand.Intn(100))
	// 	bt.Add(value, value+100000)
	// }

	var list []NodeValue
	var f1 = func(n *Node) {
		list = append(list, n.value)
	}
	bt.InOrderTraversal(f1)
	// fmt.Printf("in-order traverse binary tree -----------------------------\n%#v\n", list)

	bt.String()
	unlink := bt.Find(29)
	if unlink != nil {
		unlinked := bt.Unlink(unlink)
		fmt.Printf("SUCCESSFULLY UNLINKED NODE: %d\n", unlinked.value)
		bt.String()
		bt.InOrderTraversal(f1)
		// fmt.Printf("in-order traverse binary tree -----------------------------\n%#v\n", list)
	}

	return
}

// Item : node storage values
type Item interface{}

// NodeValue : node indexing value
type NodeValue uint64

// Node B-Tree Node Struct
type Node struct {
	value NodeValue // node indexing valuate
	item  Item      // Node storage item
	left  *Node
	right *Node
	pos   int8 //标记节点相对父节点的位置 0 tree-root, 1 left, 2 right
}

// NodeList :
type NodeList []*Node

// BTree Struct
type BTree struct {
	RootNode *Node // Root Of Binary Tree
	Len      int   // Size Of Binary Tree
}

// NewBTree Create Binary Tree
func NewBTree() (btree *BTree) {
	btree = &BTree{}
	return
}

// Add : add a new node to binary tree
func (bt *BTree) Add(value NodeValue, item Item) {
	var node Node = Node{value: value, item: item}
	if bt.RootNode == nil {
		bt.RootNode = &node
		return
	}

	if bt.RootNode.insert(&node) {
		bt.Len++
	}
}

// Insert : insert node
func (node *Node) insert(n *Node) (b bool) {
	b = false
	if n.value < node.value {
		//  把n插入node的左边
		if node.left == nil {
			n.pos = 1 // mark node as left Node
			node.left = n
		} else {
			node.left.insert(n)
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
			node.right.insert(n)
		}
		b = true
		return
	}
	b = false
	return
}

// Unlink : Remove a node from binary tree
func (bt *BTree) Unlink(un *Node) (unlink *Node) {
	if un == nil {
		return
	}
	unlink = un
	var parent *Node = bt.FindParent(unlink)
	// if unlinked & reduce length of binary tree
	defer func() {
		// 解除链接的点，左右子树指针置为 nil
		unlink.left, unlink.right = nil, nil
		// fmt.Printf("unlink-node=%d, left=%+v,right=%+v,pos=%d\n", unlink.value, unlink.left, unlink.right, unlink.pos)
		bt.Len--
	}()
	// leftTreeRoot, RightTreeRoot, unlinkedNodePosition
	left, right, pos := unlink.left, unlink.right, unlink.pos
	// fmt.Printf("parent=%+v,left=%+v,right=%+v,pos=%d\n", parent, left, right, pos)

	// relink
	var node *Node
	// ----- deal relink -------

	if left == nil && right == nil {
		// case 1: unlink node is a leaf node
		node = nil
		link(parent, node, pos)
	} else if (left != nil && right == nil) || (left == nil && right != nil) {
		// case 2: unlink node : single_child --- unlink --- parent
		node = left
		if node == nil {
			node = right
		}
		link(parent, node, pos)
	} else if left != nil && right != nil {
		// case 3: unlink node : both_child --- unlink --- parent
		// 1. try to find leftChildTree nearest node of unlink node, mark it as "nodeLN"
		nodeLN := findLeftTreeNearestNode(unlink)
		if nodeLN == left {
			left = left.left
		}
		// 2. unlink "NodeLN"
		bt.Unlink(nodeLN)
		// fmt.Printf("nodeLN = %+v\n", nodeLN)
		if parent != nil {
			// 3. unlink which node will be unlinked
			// link(parent, nil, pos) // 不是必要操作
			// 4. link parent and "NodeLN"
			link(parent, nodeLN, pos)
		}
		// 5. link "nodeLN" and LeftChildTree (left side)
		link(nodeLN, left, 1)
		// 6. link "nodeLN" and RightChildTree (right side)
		link(nodeLN, right, 2)
		node = nodeLN
	}

	// ------------------------
	if bt.RootNode == unlink {
		// 变更新 root
		bt.RootNode = node
	}
	return
}

// 从节点的左子树上找最近节点
// *查找左子树上最右节点
func findLeftTreeNearestNode(node *Node) (dist *Node) {
	if node.left == nil {
		return
	}
	// 取左树根
	dist = node.left
	for {
		// 查最右
		if dist.right == nil {
			break
		}
		dist = dist.right
	}
	return
}

func link(parent, child *Node, pos int8) {
	if pos == 1 {
		parent.left = child
	} else {
		parent.right = child
	}
	if child == nil {
		return
	}
	child.pos = pos
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

// FindParent : find parent node
func (bt *BTree) FindParent(child *Node) (parent *Node) {
	stack := bt.FindStack(child)
	if len(stack) == 0 {
		return
	}
	parent = stack[len(stack)-1:][0]
	return
}

// FindStack : find Node Stack By Child
func (bt *BTree) FindStack(child *Node) (stack NodeList) {
	visit := bt.RootNode
	for {
		if visit.value == child.value {
			return
		}
		if visit.value > child.value {
			//  当前节点比目标值大， 则查找当前节点到左树
			stack = append(stack, visit)
			visit = visit.left
		} else {
			stack = append(stack, visit)
			visit = visit.right
		}
		if visit == nil {
			return
		}
	}
}

// Find : find Node By value
func (bt *BTree) Find(value NodeValue) (node *Node) {
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
// left ->  right -> root
func (bt *BTree) PostOrderTraversal(cb func(n *Node)) {
	postOrderTraverse(bt.RootNode, cb)
}

func postOrderTraverse(node *Node, cb func(n *Node)) {
	if node != nil {
		postOrderTraverse(node.left, cb)
		postOrderTraverse(node.right, cb)
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
