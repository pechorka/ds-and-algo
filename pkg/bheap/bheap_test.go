package bheap

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryMinHeap(t *testing.T) {

	t.Run("Insert and MustPeek", func(t *testing.T) {
		h := NewMinHeap[int]()
		h.Add("5", 5)
		require.EqualValues(t, 5, h.PeekOrZero())

		h.Add("3", 3)
		require.EqualValues(t, 3, h.PeekOrZero())

		h.Add("8", 8)
		require.EqualValues(t, 3, h.PeekOrZero())

		h.Add("2", 2)
		require.EqualValues(t, 2, h.PeekOrZero())

		h.Add("4", 4)
		require.EqualValues(t, 2, h.PeekOrZero())

		h.Add("7", 7)
		require.EqualValues(t, 2, h.PeekOrZero())

		expectedHeap := []int{2, 3, 7, 5, 4, 8}
		// in bst view it is
		// 2
		// ├── 3
		// │   ├── 5
		// │   │   └── 8
		// │   └── 4
		// └── 7
		for i, v := range expectedHeap {
			require.EqualValues(t, v, h.heap[i].Val)
		}
	})

	t.Run("Extract-Min", func(t *testing.T) {
		h := NewMinHeap[int]()
		toInsert := []int{5, 3, 8, 2, 4, 7}
		for _, v := range toInsert {
			h.Add(strconv.Itoa(v), v)
		}
		val, ok := h.ExtractMin()
		require.True(t, ok)
		require.EqualValues(t, 2, val)
		require.EqualValues(t, 3, h.PeekOrZero())

		val, ok = h.ExtractMin()
		require.True(t, ok)
		require.EqualValues(t, 3, val)
		require.EqualValues(t, 4, h.PeekOrZero())

		val, ok = h.ExtractMin()
		require.True(t, ok)
		require.EqualValues(t, 4, val)
		require.EqualValues(t, 5, h.PeekOrZero())

		val, ok = h.ExtractMin()
		require.True(t, ok)
		require.EqualValues(t, 5, val)
		require.EqualValues(t, 7, h.PeekOrZero())

		val, ok = h.ExtractMin()
		require.True(t, ok)
		require.EqualValues(t, 7, val)
		require.EqualValues(t, 8, h.PeekOrZero())

		val, ok = h.ExtractMin()
		require.True(t, ok)
		require.EqualValues(t, 8, val)
	})
}
