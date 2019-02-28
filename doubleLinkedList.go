package main

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (l *LinkedList) First() *Node {
	return l.head
}

func Next(n *Node) *Node {
	return n.next
}

func (l *LinkedList) Last() *Node {
	return l.tail
}

func Previous(n *Node) *Node {
	return n.prev
}

func (l *LinkedList) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func main() {
	l := &LinkedList{}

	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	for {
		println(n.value)
		n = Next(n)
		if n == nil {
			break
		}
	}

	n = l.Last()
	for {
		println(n.value)
		n = Previous(n)
		if n == nil {
			break
		}
	}
}
