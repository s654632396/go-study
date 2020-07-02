package timewheel



type LinkNodeIface interface {
	append(newBucket LinkNodeIface)
	getID() string
	getNext() LinkNodeIface
	setNext(node LinkNodeIface)
}

type LinkNode struct {
	ID   string
	next LinkNodeIface
}

func (n *LinkNode) getID() string {
	return n.ID
}

func (n *LinkNode) append(node LinkNodeIface) {
	if n == nil {
		panic(`root LinkNode can't be nil`)
	}
	appendNode(n, node)
}
func (n *bucket) append(node LinkNodeIface) {
	if n == nil {
		panic(`root bucket can't be nil`)
	}
	appendNode(n, node)
}

func appendNode(rootNode LinkNodeIface, node LinkNodeIface) {
	var lastNode = rootNode

	var exists bool
	for {
		if lastNode.getID() == node.getID() {
			exists = true
			break
		}
		if lastNode.getNext() != nil {
			lastNode = lastNode.getNext()
		} else {
			break
		}
	}
	if exists {
		return
	}
	lastNode.setNext(node)
}

func (n *LinkNode) getNext() LinkNodeIface {
	var nb LinkNodeIface
	nb = n.next
	return nb
}

func (n *bucket) getNext() LinkNodeIface {
	var nb LinkNodeIface
	nb = n.next
	return nb
}

func (n *LinkNode) setNext(node LinkNodeIface) {
	n.next = node
}
