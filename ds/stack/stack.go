package stack

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


