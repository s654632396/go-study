package dl

import (
	"errors"
)

type DoubleList struct {
	head *ListNode
	tail *ListNode
	len  int
}

func NewDoubleList() (list *DoubleList) {
	list = &DoubleList{}
	return
}

func (list *DoubleList) Head() (node *ListNode) {
	node = list.head
	return
}

func (list *DoubleList) Tail() (node *ListNode) {
	node = list.tail
	return
}

func (list *DoubleList) Len() (len int) {
	len = list.len
	return
}

func (list *DoubleList) LPop() (node *ListNode) {
	if list.len == 0 {
		return
	}
	node = list.head

	if node.next != nil {
		list.head = node.next
		list.head.prev = nil
	} else {
		list.head = nil
		list.tail = nil
	}
	list.len--
	return
}

func (list *DoubleList) RPush(node_value string) {
	node := NewListNode(node_value)

	if list.len == 0 {
		list.head = node
		list.tail = node
	} else {
		tail := list.tail
		tail.next = node
		node.prev = tail

		list.tail = node
	}
	list.len++
	return
}

func (list *DoubleList) Index(index int) (node *ListNode) {

	if index < 0 {

		index = -index - 1
		if index > list.len {
			return
		}
		// 从尾部逆向查找
		node = list.tail
		for ; index > 0 && node != nil; index-- {
			node = node.Prev()
		}
	} else {
		if index > list.len+1 {
			return
		}
		node = list.head
		// println("[debug]: ", node.Next().Next().Next().value)

		for ; index > 0 && node != nil; index-- {
			node = node.Next()
		}
	}
	return
}

func (list *DoubleList) Range(start, len int) (nodes []*ListNode, err error) {
	nodes = make([]*ListNode, 0)

	if len <= 0 {
		err = errors.New("len must larger than 0.")
		return
	}

	var node *ListNode
	for i := 0; i < len; i++ {
		if i == 0 {
			node = list.Index(start)
		} else {
			node = node.Next()
		}
		if node == nil {
			break
		}
		nodes = append(nodes, node)
	}
	return
}
