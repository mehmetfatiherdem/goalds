package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type Stack[T comparable] struct {
	Head *Node[T]
}

func (s *Stack[T]) Search(n T) (int, error) {
	index := 1
	it := s.Head
	for it.Value != n {
		it = it.Next
		index++
	}

	if it == nil {
		panic("value not found in the stack")
	}

	return index, nil
} 

func (s *Stack[T]) Empty() bool {
	return s.Head == nil
}

func (s *Stack[T]) Peek() (n *Node[T]) {
	return s.Head
}

func (s *Stack[T]) Push(val T) {
	n := Node[T]{Value: val, Next: nil}
	n.Next = s.Head
	s.Head = &n
}

func (s *Stack[T]) Pop() *Node[T] {
	poppedHead := s.Head
	s.Head = s.Head.Next
	return poppedHead
}

func main() {
	stack := Stack[int]{}

	fmt.Println("stack is empty before:", stack.Empty())
	
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	fmt.Println("stack is empty after:", stack.Empty())

	for stack.Head != nil {
		fmt.Println(stack.Head.Value)
		stack.Pop()
	}

	fmt.Println("stack is empty after print before search:", stack.Empty())

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)
	
	in, err := stack.Search(40)

	if err != nil {
		panic(err)
	}

	fmt.Println("searching index of 40:", in)
	
}
