package bst

type BST[v comparable] interface {
	Delete(Val v)

	Search(Val v) *Bstnode[v]

	Insert(Val v)

	Traversal() []v
}
