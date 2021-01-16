package main

import "fmt"

type LinkList struct {
	Value interface{}
	next  *LinkList
}

func NewLinkListNode(value interface{}) *LinkList {
	node := &LinkList{}
	node.Value = value
	node.next = nil
	return node
}

func (ll *LinkList) After(next *LinkList) {
	var curNode = ll
	for {
		if curNode.next == nil {
			curNode.next = next
			break
		} else {
			curNode = curNode.next
		}
	}
}

func DumpLinkList (rootNode *LinkList) {
	var curNode = rootNode
	for {
		fmt.Print(curNode.Value , "-")
		if curNode.next == nil {
			break
		}
		curNode = curNode.next
	}
}

func main() {
	dataList := []int{4, 52, 32, 67, 31, 67}

	var root *LinkList
	for _, data := range dataList {
		if root != nil {
			node := NewLinkListNode(data)
			root.After(node)
		} else {
			// init
			root = NewLinkListNode(data)
		}
	}


	DumpLinkList(root)
	fmt.Println()
	reverseLinkList(root)
}

func reverseLinkList(rootNode *LinkList)  {
	recursivePrintNode(rootNode)
}

func recursivePrintNode(node *LinkList) {
	if node.next == nil {
		fmt.Print(node.Value, "-")
	} else {
		recursivePrintNode(node.next)
		fmt.Print(node.Value, "-")
	}
}