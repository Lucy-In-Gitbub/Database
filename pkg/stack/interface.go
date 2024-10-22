package stack

type Stackdata interface {
	IsEmpty() bool
	Push(value any)
	Pop() any
	Size() int
}
