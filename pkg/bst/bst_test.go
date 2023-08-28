package bst

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBST(t *testing.T) {
	t.Run("Insert and Search", func(t *testing.T) {
		prepareTree := func() *BST {
			bst := &BST{}
			values := []int{5, 3, 8, 1, 4, 7, 9}

			for _, v := range values {
				bst.Insert(v)
			}
			return bst
		}

		t.Run("happy case", func(t *testing.T) {
			bst := prepareTree()
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

		t.Run("insert duplicate", func(t *testing.T) {
			bst := New()
			bst.Insert(3)
			bst.Insert(3)
			require.NotNil(t, bst.root)
			require.Nil(t, bst.root.left)
			require.Nil(t, bst.root.right)
		})
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

	t.Run("Traverse", func(t *testing.T) {
		prepareTree := func() *BST {
			bst := New()
			values := []int{50, 30, 70, 20, 40, 60, 80}
			for _, v := range values {
				bst.Insert(v)
			}
			// 50
			// ├── 30
			// │   ├── 20
			// │   └── 40
			// └── 70
			//     ├── 60
			//     └── 80
			return bst
		}

		t.Run("InOrder", func(t *testing.T) {
			bst := prepareTree()
			expected := []int{20, 30, 40, 50, 60, 70, 80}
			actual := bst.TraverseInOrder()
			require.Equal(t, expected, actual)
		})

		t.Run("PreOrder", func(t *testing.T) {
			bst := prepareTree()
			expected := []int{50, 30, 20, 40, 70, 60, 80}
			actual := bst.TraversePreOrder()
			require.Equal(t, expected, actual)
		})

		t.Run("PostOrder", func(t *testing.T) {
			bst := prepareTree()
			expected := []int{20, 40, 30, 60, 80, 70, 50}
			actual := bst.TraversePostOrder()
			require.Equal(t, expected, actual)
		})

		t.Run("LevelOrder", func(t *testing.T) {
			bst := prepareTree()
			expected := []int{50, 30, 70, 20, 40, 60, 80}
			actual := bst.TraverseLevelOrder()
			require.Equal(t, expected, actual)
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

func TestDrawTree(t *testing.T) {
	t.Run("Draw larger tree", func(t *testing.T) {
		bst := New()
		values := []int{50, 30, 70, 20, 40, 60, 80}
		for _, v := range values {
			bst.Insert(v)
		}

		expectedDrawing := strings.Join([]string{
			"50",
			"├── 30",
			"│   ├── 20",
			"│   └── 40",
			"└── 70",
			"    ├── 60",
			"    └── 80",
		}, "\n")
		expectedDrawing += "\n"

		tree := bst.DrawTree()

		require.Equal(t, expectedDrawing, tree)
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
