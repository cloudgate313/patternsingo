package main

type BinarySearchTree struct {
	node *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (t *BinarySearchTree) TreeInsert(value int) *BinarySearchTree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.NodeInsert(value)
	}

	return t
}

func (n *Node) NodeInsert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.NodeInsert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.NodeInsert(value)
		}
	}
}

func PrintNode(n *Node) {
	if n == nil {
		return
	}

	println(n.value)
	PrintNode(n.left)
	PrintNode(n.right)
}

func main() {
	t := &BinarySearchTree{}
	t.TreeInsert(1).
		TreeInsert(2).
		TreeInsert(3).
		TreeInsert(4).
		TreeInsert(5)

	PrintNode(t.node)
}
