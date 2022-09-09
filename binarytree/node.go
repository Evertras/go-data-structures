package avltree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	size int
	val  T

	left  *Node[T]
	right *Node[T]
}

func newNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{
		size: 1,
		val:  val,
	}
}

func (n *Node[T]) Value() T {
	return n.val
}

func (n *Node[T]) Size() int {
	return n.size
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func (n *Node[T]) insert(val T) {
	n.size++

	if val < n.val {
		if n.left == nil {
			n.left = newNode(val)
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = newNode(val)
		} else {
			n.right.insert(val)
		}
	}
}
