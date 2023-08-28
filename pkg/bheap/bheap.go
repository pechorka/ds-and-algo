package bheap

import "cmp"

type value[V cmp.Ordered] struct {
	ID  string
	Val V
}

type MinBinaryHeap[V cmp.Ordered] struct {
	// use a slice to store the heap (binary tree)
	// the first element is the root
	// the children of the element at index i are at indices 2i+1 (left) and 2i+2 (right)
	heap    []value[V]
	idToIdx map[string]int // map from ID to index in heap
}

func NewMinHeap[V cmp.Ordered]() *MinBinaryHeap[V] {
	return &MinBinaryHeap[V]{
		idToIdx: make(map[string]int),
	}
}

func (h *MinBinaryHeap[V]) Add(id string, val V) bool {
	if _, ok := h.idToIdx[id]; ok {
		return false
	}
	h.heap = append(h.heap, value[V]{
		ID:  id,
		Val: val,
	})
	h.idToIdx[id] = len(h.heap) - 1
	h.bubbleUp()
	return true
}

// bubbleUp rebalances the heap from the bottom up
func (h *MinBinaryHeap[V]) bubbleUp() {
	i := len(h.heap) - 1
	for {
		parent := (i - 1) / 2
		if parent == i || h.heap[i].Val >= h.heap[parent].Val {
			break
		}
		h.swap(i, parent)
		i = parent
	}
}

func (h *MinBinaryHeap[V]) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
	// update index maps
	h.idToIdx[h.heap[i].ID] = i
	h.idToIdx[h.heap[j].ID] = j
}

func (h *MinBinaryHeap[V]) ExtractMin() (id string, val V, ok bool) {
	if len(h.heap) == 0 {
		return "", val, false
	}
	internalVal := h.heap[0]
	h.swap(0, len(h.heap)-1)
	h.heap = h.heap[:len(h.heap)-1]
	delete(h.idToIdx, internalVal.ID)
	h.bubbleDown()
	return internalVal.ID, internalVal.Val, true
}

// bubbleDown rebalances the heap from the top down
func (h *MinBinaryHeap[V]) bubbleDown() {
	i := 0
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < len(h.heap) && h.heap[left].Val < h.heap[smallest].Val {
			smallest = left
		}
		if right < len(h.heap) && h.heap[right].Val < h.heap[smallest].Val {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.swap(i, smallest)
		i = smallest
	}
}

func (h *MinBinaryHeap[V]) Peek() (val V, ok bool) {
	if len(h.heap) == 0 {
		return val, false
	}
	return h.heap[0].Val, true
}

func (h *MinBinaryHeap[V]) PeekOrZero() (val V) {
	if len(h.heap) == 0 {
		return val
	}
	return h.heap[0].Val
}

func (h *MinBinaryHeap[V]) Update(id string, newVal V) {
	idx, ok := h.idToIdx[id]
	if !ok {
		return
	}

	curVal := h.heap[idx]
	if curVal.Val == newVal {
		return
	}

	h.heap[idx] = value[V]{
		ID:  id,
		Val: newVal,
	}
	if curVal.Val > newVal {
		h.bubbleUp()
	} else {
		h.bubbleDown()
	}
}
