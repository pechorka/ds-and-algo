package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBST(t *testing.T) {
	t.Run("Insert and Search", func(t *testing.T) {
		bst := &BST{}
		values := []int{5, 3, 8, 1, 4, 7, 9}

		for _, v := range values {
			bst.Insert(v)
		}

		require.True(t, bst.Search(5))
		require.True(t, bst.Search(3))
		require.True(t, bst.Search(8))
		require.True(t, bst.Search(1))
		require.True(t, bst.Search(4))
		require.True(t, bst.Search(7))
		require.True(t, bst.Search(9))
		require.False(t, bst.Search(0))
		require.False(t, bst.Search(6))
		require.False(t, bst.Search(10))
	})

	t.Run("Delete", func(t *testing.T) {
		prepareTree := func() *BST {
			bst := New()
			values := []int{50, 30, 70, 20, 40, 60, 80}
			for _, v := range values {
				bst.Insert(v)
			}
			return bst
		}

		t.Run("Delete leaf node", func(t *testing.T) {
			bst := prepareTree()
			bst.Delete(20)
			require.False(t, bst.Search(20))
		})

		t.Run("Delete node with one child", func(t *testing.T) {
			bst := prepareTree()
			bst.Delete(30)
			require.False(t, bst.Search(30))
			require.True(t, bst.Search(40))
		})

		t.Run("Delete node with two children", func(t *testing.T) {
			bst := prepareTree()
			bst.Delete(50)
			require.False(t, bst.Search(50))
			// Check that children are still there
			require.True(t, bst.Search(60) && bst.Search(40))
		})
	})

	t.Run("BST Properties", func(t *testing.T) {
		bst := &BST{}
		values := []int{5, 3, 8, 1, 4, 7, 9}

		for _, v := range values {
			bst.Insert(v)
		}

		// Assuming you have a method to get the left child and right child
		// rootValue should be 5
		rootValue := bst.root.value
		require.Equal(t, 5, rootValue)

		// Left child of root should be 3
		leftChildValue := bst.root.left.value
		require.Equal(t, 3, leftChildValue)

		// Right child of root should be 8
		rightChildValue := bst.root.right.value
		require.Equal(t, 8, rightChildValue)
	})
}

func BenchmarkInsert(b *testing.B) {
	bst := &BST{}
	for i := 0; i < b.N; i++ {
		bst.Insert(i)
	}
}

func BenchmarkSearch(b *testing.B) {
	bst := &BST{}
	for i := 0; i < 10000; i++ {
		bst.Insert(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bst.Search(i % 10000)
	}
}
