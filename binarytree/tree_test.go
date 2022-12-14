package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	tree := New[int]()

	assert.Equal(t, 0, tree.Size())

	assert.Nil(t, tree.Root())
	assert.Nil(t, tree.Root().Parent())
}

func TestInsertRoot(t *testing.T) {
	tree := New[int]()

	const insertedVal = 3

	tree.Insert(insertedVal)

	assert.Equal(t, 1, tree.Size())

	root := tree.Root()

	assert.Equal(t, insertedVal, root.Value())
	assert.Nil(t, tree.Root().Parent())
}

func TestInsertSmallerValue(t *testing.T) {
	tree := New[int]()

	const (
		rootVal     = 3
		insertedVal = 2
	)

	tree.Insert(rootVal)
	tree.Insert(insertedVal)

	assert.Equal(t, 2, tree.Size())
	assert.Equal(t, 2, tree.Height())

	assert.Equal(t, rootVal, tree.Root().Value())
	assert.NotNil(t, tree.Root().Left())
	assert.Nil(t, tree.Root().Right())
	assert.Equal(t, insertedVal, tree.Root().Left().Value())
	assert.Equal(t, tree.Root(), tree.Root().Left().Parent())
}

func TestInsertLargerValue(t *testing.T) {
	tree := New[int]()

	const (
		rootVal     = 3
		insertedVal = 5
	)

	tree.Insert(rootVal)
	tree.Insert(insertedVal)

	assert.Equal(t, 2, tree.Size())
	assert.Equal(t, 2, tree.Height())

	assert.Equal(t, rootVal, tree.Root().Value())
	assert.NotNil(t, tree.Root().Right())
	assert.Nil(t, tree.Root().Left())
	assert.Equal(t, insertedVal, tree.Root().Right().Value())
}

func TestInsertBothSmallerAndLarger(t *testing.T) {
	tree := New[int]()

	const (
		rootVal    = 3
		smallerVal = 1
		largerVal  = 5
	)

	tree.Insert(rootVal)
	tree.Insert(smallerVal)
	tree.Insert(largerVal)

	assert.Equal(t, 3, tree.Size())
	assert.Equal(t, 2, tree.Height())

	assert.Equal(t, rootVal, tree.Root().Value())
	assert.NotNil(t, tree.Root().Right())
	assert.NotNil(t, tree.Root().Left())
	assert.Equal(t, largerVal, tree.Root().Right().Value())
	assert.Equal(t, smallerVal, tree.Root().Left().Value())
}

func TestMultipleLevels(t *testing.T) {
	tree := New[int]()

	const (
		rootVal    = 3
		smallerVal = 1
		midVal     = 2
	)

	tree.Insert(rootVal)
	tree.Insert(smallerVal)
	tree.Insert(midVal)

	assert.Equal(t, 3, tree.Size())
	assert.Equal(t, 3, tree.Height())

	assert.Equal(t, rootVal, tree.Root().Value())
	assert.NotNil(t, tree.Root().Left())
	assert.NotNil(t, tree.Root().Left().Right())
	assert.Nil(t, tree.Root().Right())
	assert.Equal(t, smallerVal, tree.Root().Left().Value())
	assert.Equal(t, midVal, tree.Root().Left().Right().Value())

	assert.Equal(t, 2, tree.Root().Left().Size())
}
