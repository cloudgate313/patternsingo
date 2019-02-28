package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	last := len(s.items)

	if last == 0 {
		return -1
	}

	item, items := s.items[last-1], s.items[0:last-1]
	s.items = items
	return item
}

func main() {
	s := &Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s)

	println(s.Pop())
	println(s.Pop())
	fmt.Println(s)

	println(s.Pop())
	println(s.Pop())

	fmt.Println(s)

}
