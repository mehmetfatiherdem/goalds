package main

import "fmt"

type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}

type Stack[T interface{}] struct {
	Head *Node[T]
}

func (s *Stack[T]) Push(n *Node[T]) {
	n.Next = s.Head
	s.Head = n
}

func (s *Stack[T]) Pop() *Node[T] {
	poppedHead := s.Head
	s.Head = s.Head.Next
	return poppedHead
}

func main() {
	head := Node[int]{Value: 10, Next: nil}

	stack := Stack[int]{Head: &head}

	stack.Push(&Node[int]{Value: 20, Next: nil})
	stack.Push(&Node[int]{Value: 30, Next: nil})
	stack.Push(&Node[int]{Value: 40, Next: nil})

	for stack.Head != nil {
		fmt.Println(stack.Head.Value)
		stack.Pop()
	}
}
