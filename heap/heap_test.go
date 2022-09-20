package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapEmpty(t *testing.T) {
	h := NewMinHeap[int]()

	assert.True(t, h.IsEmpty())

	h.Insert(3)

	assert.False(t, h.IsEmpty())
}

func TestHeapInsert(t *testing.T) {
	h := NewMinHeap[int]()

	assert.True(t, h.IsEmpty())

	h.Insert(3)

	assert.False(t, h.IsEmpty())
	assert.Equal(t, 3, h.Top())

	h.Insert(5)

	assert.Equal(t, 3, h.Top())

	h.Insert(2)

	assert.Equal(t, 2, h.Top())

	for i := 3; i < 100; i++ {
		h.Insert(i)
	}

	assert.Equal(t, 2, h.Top())

	h.Insert(1)

	assert.Equal(t, 1, h.Top())
}

func TestHeapPop(t *testing.T) {
	h := NewMinHeap[int]()

	h.Insert(3)

	assert.False(t, h.IsEmpty())

	popped := h.Pop()

	assert.Equal(t, 3, popped)
	assert.True(t, h.IsEmpty())

	for i := 100; i > 0; i-- {
		h.Insert(i)
	}

	// Check twice to make sure unmodified
	assert.Equal(t, 1, h.Top())
	assert.Equal(t, 1, h.Top())

	// Check multiple successive operations
	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 2, h.Pop())
	assert.Equal(t, 3, h.Pop())
	assert.Equal(t, 4, h.Pop())
}

func BenchmarkHeapify10k(b *testing.B) {
	h := NewMinHeap[int]()

	for i := 1; i <= 10000; i++ {
		h.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		h.heapify(0)
	}
}

func BenchmarkInsertInto10k(b *testing.B) {
	h := NewMinHeap[int]()

	for i := 1; i <= 10000; i++ {
		h.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		h.Insert(0)
		// Reset
		h.vals[0] = 1
		h.vals = h.vals[:10000]
	}
}

func BenchmarkPopFrom10k(b *testing.B) {
	h := NewMinHeap[int]()

	for i := 1; i <= 10000; i++ {
		h.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		h.Pop()
		// Reset, we can subtract this later from benchmark above for approximate
		h.Insert(0)
	}
}
