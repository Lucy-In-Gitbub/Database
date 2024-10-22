package list

import (
	"errors"
)

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
	Prev *ListNode[T]
}

type List[T comparable] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
	Hash map[T][]*ListNode[T]
}

type Iterator[T comparable] struct {
	Pointer *ListNode[T]
}

func Create[T comparable](Val T) *List[T] {
	head := &ListNode[T]{Prev: nil}
	next := &ListNode[T]{Val: Val, Next: nil}
	head.Next = next
	next.Prev = head

	hash := map[T][]*ListNode[T]{}
	hash[Val] = append(hash[Val], next)

	return &List[T]{head, next, hash}
}

func (l *List[T]) Find(Val T) ([]*ListNode[T], bool) {
	p, ok := l.Hash[Val]

	return p, ok
}

// head insert
func (l *List[T]) Add(Val T) {
	new := &ListNode[T]{Val: Val, Next: l.Head.Next}
	next := l.Head.Next
	next.Prev = new
	l.Head.Next = new

	l.Hash[Val] = append(l.Hash[Val], new)
}

func (l *List[T]) Delete(Val T, index int) error {
	p, ok := l.Find(Val)

	if ok {
		if index < 0 || index >= len(p) {
			return errors.New("index out of range")
		} else {
			// delete node
			del := l.Hash[Val][index]
			prev := del.Prev
			next := del.Next
			prev.Next = next
			next.Prev = prev

			//delete hash
			l.Hash[Val] = append(l.Hash[Val][0:index], l.Hash[Val][index+1:]...)
		}
	} else {
		return errors.New("not found")
	}

	return nil
}

func (l *List[T]) Begin() Iterator[T] {
	var Iter Iterator[T]
	Iter.Pointer = l.Head.Next

	return Iter
}

func (l *List[T]) End() Iterator[T] {
	var Iter Iterator[T]
	Iter.Pointer = l.Tail

	return Iter
}

func (Iter Iterator[T]) Get() T {
	return Iter.Pointer.Val
}

func (Iter Iterator[T]) IsEmpty() bool {
	return Iter.Pointer == nil
}

func (Iter Iterator[T]) Next() Iterator[T] {
	Iter.Pointer = Iter.Pointer.Next

	return Iter
}

func (Iter Iterator[T]) Prev() Iterator[T] {
	Iter.Pointer = Iter.Pointer.Prev

	return Iter
}
