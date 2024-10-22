package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := Create[int]()
	s.Push(10)
	n := s.Pop()
	fmt.Println(n)
	fmt.Println(s.Size())
	b := s.IsEmpty()
	fmt.Println(b)
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Size())
	fmt.Println(s.Peek())
	fmt.Println(s.IsEmpty())
	s.Pop()
	n = s.Pop()
	fmt.Println(n)
}
