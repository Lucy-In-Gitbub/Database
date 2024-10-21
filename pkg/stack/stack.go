package stack

import "fmt"

type Node[T any] struct {
	Value T
	Blow  *Node[T]
}

type Stack[T any] struct {
	Top  *Node[T]
	size int
}

// information, error handling
func handPanic() {
	if s := recover(); s != nil {
		fmt.Println("panic:", s)
	}
}

func (s *Stack[T]) Size() int {

	return s.size
}

func (s *Stack[T]) IsEmpty() bool {

	return s.Top == nil
}

func (s *Stack[T]) Peek() T {
	//Empty check
	if s.IsEmpty() {
		defer handPanic()
		panic("stack is empty")
	}

	return s.Top.Value
}

// function
func Create[T any]() *Stack[T] {
	stack := &Stack[T]{nil, 0}

	return stack
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{value, s.Top}
	s.Top = node

	s.size++
}

func (s *Stack[T]) Pop() T {
	//Empty check
	if s.IsEmpty() {
		defer handPanic()
		panic("stack is empty")
	}

	value := s.Top.Value
	s.Top = s.Top.Blow

	s.size--
	return value
}
