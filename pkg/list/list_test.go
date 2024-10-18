package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list1 := Create("apple")
	list1.Add("banana")
	list1.Add("orange")
	list1.Add("apple")
	x, ok := list1.Find("apple")
	if ok {
		fmt.Println("Found apple", x)
	} else {
		fmt.Println("Not found apple")
	}
	list1.Delete("banana", 0)
	Iter := list1.Begin()
	for !Iter.IsEmpty() {
		fmt.Println(Iter.Pointer.Val)
		Iter = Iter.Next()
	}
	list1.Add("banana")
	err := list1.Delete("grape", 0)
	fmt.Println(err)
	err = list1.Delete("apple", 2)
	fmt.Println(err)
	list1.Delete("apple", 1)
	Iter = list1.Begin()
	for !Iter.IsEmpty() {
		fmt.Println(Iter.Pointer.Val)
		Iter = Iter.Next()
	}
}
