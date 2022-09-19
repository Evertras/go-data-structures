package heap

import "golang.org/x/exp/constraints"

type Heap[T interface{}] struct {
	vals        []T
	compareFunc func(a, b T) bool
}

func NewHeap[T interface{}](compareFunc func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		vals:        []T{},
		compareFunc: compareFunc,
	}
}

func NewMinHeap[T constraints.Ordered](initialVals ...T) *Heap[T] {
	h := NewHeap(func(a, b T) bool { return a < b })

	for _, v := range initialVals {
		h.Insert(v)
	}

	return h
}

func NewMaxHeap[T constraints.Ordered](initialVals ...T) *Heap[T] {
	h := NewHeap(func(a, b T) bool { return a > b })

	for _, v := range initialVals {
		h.Insert(v)
	}

	return h
}

func (h *Heap[T]) Insert(a T) *Heap[T] {
	h.vals = append(h.vals, a)
	l := len(h.vals)

	for iParent := l - 1; iParent >= 0; iParent = (iParent - 1) / 2 {
		iLeft := iParent*2 + 1
		iRight := iParent*2 + 2
		iLargest := iParent

		if iLeft < l && h.compareFunc(h.vals[iLeft], h.vals[iLargest]) {
			iLargest = iLeft
		}

		if iRight < l && h.compareFunc(h.vals[iRight], h.vals[iLargest]) {
			iLargest = iRight
		}

		if iLargest != iParent {
			tmp := h.vals[iLargest]
			h.vals[iLargest] = h.vals[iParent]
			h.vals[iParent] = tmp
		}

		if iParent == 0 {
			return h
		}
	}

	return h
}

func (h *Heap[T]) Top() T {
	if len(h.vals) == 0 {
		// A bit scary but WAY cleaner API this way for testing/funsies
		panic("heap is empty")
	}

	return h.vals[0]
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.vals) == 0
}

func (h *Heap[T]) heapify(iRoot int) *Heap[T] {
	l := len(h.vals)
	iLargest := iRoot
	iLeft := iRoot*2 + 1
	iRight := iRoot*2 + 2

	if iLeft < l {
		h.heapify(iLeft)
	}

	if iRight < l {
		h.heapify(iRight)
	}

	if iLeft < l &&
		h.compareFunc(
			h.vals[iLeft],
			h.vals[iLargest],
		) {
		iLargest = iLeft
	}

	if iRight < l &&
		h.compareFunc(
			h.vals[iRight],
			h.vals[iLargest],
		) {
		iLargest = iRight
	}

	if iLargest != iRoot {
		tmp := h.vals[iRoot]
		h.vals[iRoot] = h.vals[iLargest]
		h.vals[iLargest] = tmp
	}

	return h
}
