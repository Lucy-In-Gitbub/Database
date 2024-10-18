package bst

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	bst := Create[int]()
	//	bst.Delete(1)
	//	bst.Search(1)
	bst.Insert(7)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(10)
	bst.Insert(12)
	bst.Insert(11)

	res := bst.Traversal()
	for x, v := range res {
		fmt.Printf("%d: %d\n", x, v)
	}

	bst.Delete(12)
	res = bst.Traversal()
	for x, v := range res {
		fmt.Printf("%d: %d\n", x, v)
	}
	bst.Delete(10)
	res = bst.Traversal()
	for x, v := range res {
		fmt.Printf("%d: %d\n", x, v)
	}
	bst.Delete(1)
	res = bst.Traversal()
	for x, v := range res {
		fmt.Printf("%d: %d\n", x, v)
	}
	bst.Delete(3)
	res = bst.Traversal()
	for x, v := range res {
		fmt.Printf("%d: %d\n", x, v)
	}
}