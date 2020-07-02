package dl

type ListNode struct {
	value string
	prev  *ListNode
	next  *ListNode
}

func NewListNode(value string) (listNode *ListNode) {
	return &ListNode{
		value: value,
	}
}

func (node *ListNode) Prev() (prev *ListNode) {
	prev = node.prev
	return
}

func (node *ListNode) Next() (next *ListNode) {
	next = node.next
	return
}

func (node *ListNode) Value() (value string) {
	value = node.value
	return
}
