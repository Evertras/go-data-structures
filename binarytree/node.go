package binarytree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	size   int
	height int
	val    T

	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

func newNode[T constraints.Ordered](parent *Node[T], val T) *Node[T] {
	return &Node[T]{
		parent: parent,
		size:   1,
		height: 1,
		val:    val,
	}
}

func (n *Node[T]) Parent() *Node[T] {
	if n == nil {
		return nil
	}

	return n.parent
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
	if n == nil {
		return nil
	}

	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	if n == nil {
		return nil
	}

	return n.right
}

func (n *Node[T]) insert(val T) (increasedHeight bool) {
	n.size++

	if val < n.val {
		if n.left == nil {
			n.left = newNode(n, val)
			increasedHeight = n.right == nil
		} else {
			increasedHeight = n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = newNode(n, val)
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
