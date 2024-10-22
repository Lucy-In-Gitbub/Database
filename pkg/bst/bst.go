package bst

import (
	"cmp"
	"fmt"

	"github.com/Lucy-In-Gitbub/Datastruct/pkg/stack"
)

type Comparator[T any] func(x, y T) int

type Bstnode[v comparable] struct {
	Val   v
	Left  *Bstnode[v]
	Right *Bstnode[v]
	Comp  Comparator[v]
}

func NewNode[v cmp.Ordered]() *Bstnode[v] {
	return &Bstnode[v]{Comp: cmp.Compare[v]}
}

type Bst[v cmp.Ordered] struct {
	root *Bstnode[v]
}

// function
func Create[v cmp.Ordered]() *Bst[v] {
	return &Bst[v]{root: nil}
}

func (bst *Bst[v]) Insert(Val v) {
	var pointer **Bstnode[v] = &bst.root

	//if tree is empty
	if IsEmpty(*pointer) {
		*pointer = NewNode[v]()
		(*pointer).Val = Val
		return
	}

	for !IsEmpty(*pointer) {
		switch cmp.Compare(Val, (*pointer).Val) {
		case -1:
			fallthrough
		case 0:
			pointer = &((*pointer).Left)
		case 1:
			pointer = &((*pointer).Right)
		}
	}

	*pointer = NewNode[v]()
	(*pointer).Val = Val
}

func (bst *Bst[v]) Delete(Val v) {
	pointer := bst.root

	//if tree is empty
	if IsEmpty(pointer) {
		defer handPanic()
		panic("tree is empty")
	}

	var parent *Bstnode[v] = nil
	for !IsEmpty(pointer) {
		switch cmp.Compare(Val, pointer.Val) {
		case -1:
			parent = pointer
			pointer = pointer.Left
		case 0:
			//r,l is empty
			if IsEmpty(pointer.Left) && IsEmpty(pointer.Right) {
				switch cmp.Compare(parent.Val, pointer.Val) {
				case -1:
					fallthrough
				case 0:
					parent.Left = nil
				case 1:
					parent.Right = nil
				}
				return
			}

			//r is emypty,l is not empty
			if IsEmpty(pointer.Right) && !IsEmpty(pointer.Left) {
				switch cmp.Compare(pointer.Val, parent.Val) {
				case -1:
					fallthrough
				case 0:
					parent.Left = pointer.Left
				case 1:
					parent.Right = pointer.Left
				}
				return
			}

			//l is empty,r is not empty
			if !IsEmpty(pointer.Right) && IsEmpty(pointer.Left) {
				switch cmp.Compare(pointer.Val, parent.Val) {
				case -1:
					fallthrough
				case 0:
					parent.Left = pointer.Right
				case 1:
					parent.Right = pointer.Right
				}
				return
			}

			//r,l is not empty
			r := pointer.Right
			var pr *Bstnode[v] = nil
			for !IsEmpty(r) {
				pr = r
				r = r.Left
			}
			pr.Left = pointer.Left
			parent.Left = pointer.Right
			pointer.Left = nil
			pointer.Right = nil

			return

		case 1:
			parent = pointer
			pointer = pointer.Right
		}
	}
}

func (bst *Bst[v]) Search(Val v) *Bstnode[v] {
	pointer := bst.root

	//if tree is empty
	if IsEmpty(pointer) {
		defer handPanic()
		panic("tree is empty")
	}

	for !IsEmpty(pointer) {
		switch cmp.Compare(Val, pointer.Val) {
		case -1:
			pointer = pointer.Left
		case 0:
			return pointer
		case 1:
			pointer = pointer.Right
		}
	}

	defer handPanic()
	panic("not found")
}

// information
func IsEmpty[v comparable](n *Bstnode[v]) bool {
	return n == nil
}

func handPanic() {
	if s := recover(); s != nil {
		fmt.Println("panic:", s)
	}
}

func (bst *Bst[v]) Traversal() []v {
	if IsEmpty(bst.root) {
		return []v{}
	}

	stack := stack.Create[*Bstnode[v]]()
	res := []v{}

	pointer := bst.root
	for stack.Size() > 0 || !IsEmpty(pointer) {
		if !IsEmpty(pointer) {
			stack.Push(pointer)
			pointer = pointer.Left
		} else {
			cur := stack.Pop()
			pointer = cur.Right
			res = append(res, cur.Val)
		}
	}

	return res
}

func (bst *Bst[v]) Dfs() []v {
	if IsEmpty(bst.root) {
		return []v{}
	}

	stack := stack.Create[*Bstnode[v]]()
	res := []v{}

	pointer := bst.root
	for stack.Size() > 0 || !IsEmpty(pointer) {
		if !IsEmpty(pointer) {
			res = append(res, pointer.Val)
			stack.Push(pointer)
			pointer = pointer.Left
		} else {
			cur := stack.Pop()
			pointer = cur.Right
		}
	}

	return res
}

func (bst *Bst[v]) Bfs() []v {
	if IsEmpty(bst.root) {
		return []v{}
	}

	s1 := stack.Create[*Bstnode[v]]()
	s2 := stack.Create[*Bstnode[v]]()
	res := []v{}

	pointer := bst.root
	for !IsEmpty(pointer) || s1.Size() > 0 {
		if !IsEmpty(pointer) {
			s1.Push(pointer)
			s2.Push(pointer)
			pointer = pointer.Right
		} else {
			pointer = s1.Pop()
			pointer = pointer.Left
		}
	}

	for s2.Size() > 0 {
		pointer = s2.Pop()
		res = append(res, pointer.Val)
	}

	return res
}
