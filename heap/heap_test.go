package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestHeapEmpty(t *testing.T) {
	h := NewMinHeap[int]()

	assert.True(t, h.IsEmpty())

	h.Insert(3)

	assert.False(t, h.IsEmpty())
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

func BenchmarkInsert10k(b *testing.B) {
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
