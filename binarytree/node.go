package avltree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	size   int
	height int
	val    T

	left  *Node[T]
	right *Node[T]
}

func newNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{
		size:   1,
		height: 1,
		val:    val,
	}
}

func (n *Node[T]) Value() T {
	return n.val
}

func (n *Node[T]) Size() int {
	if n == nil {
		return 0
	}

	return n.size
}

func (n *Node[T]) Height() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func (n *Node[T]) insert(val T) (increasedHeight bool) {
	n.size++

	if val < n.val {
		if n.left == nil {
			n.left = newNode(val)
			increasedHeight = n.right == nil
		} else {
			increasedHeight = n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = newNode(val)
			increasedHeight = n.left == nil
		} else {
			increasedHeight = n.right.insert(val)
		}
	}

	if increasedHeight {
		n.height++
	}

	return
}
