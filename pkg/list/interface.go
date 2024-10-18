package list

type LinkedList[T comparable] interface {
	Find(Val T) (*ListNode[T], bool)
	// find a node with value Val in the list, return the node and a bool to indicate whether it was found or not

	Add(Val T)
	// add a new node with value Val to the head of the list

	Delete(Val T, index int) error

	Begin() Iterator[T]

	End() Iterator[T]
}

type Iterat[T comparable] interface {
	Begin() Iterator[T]

	End() Iterator[T]

	Next() Iterator[T]

	Prev() Iterator[T]

	Get() T

	IsEmpty() bool
}
