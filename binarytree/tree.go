package avltree

import "golang.org/x/exp/constraints"

type Tree[T constraints.Ordered] struct {
	root *Node[T]
}

func (t *Tree[T]) Size() int {
	if t.root == nil {
		return 0
	}

	return t.root.Size()
}

func (t *Tree[T]) Height() int {
	return t.root.Height()
}

func New[T constraints.Ordered]() Tree[T] {
	return Tree[T]{}
}

func (t *Tree[T]) Insert(val T) {
	if t.root == nil {
		t.root = newNode(val)
		return
	}

	t.root.insert(val)
}

func (t *Tree[T]) Root() *Node[T] {
	return t.root
}
