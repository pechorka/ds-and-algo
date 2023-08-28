package bst

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBST(t *testing.T) {
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
	t.Run("Insert and Search", func(t *testing.T) {
		t.Run("happy case", func(t *testing.T) {
			values := []int{5, 3, 8, 1, 4, 7, 9}
			bst := New()
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

	t.Run("Height", func(t *testing.T) {
		t.Run("Height of the tree", func(t *testing.T) {
			bst := prepareTree()
			require.Equal(t, 3, bst.Height())
		})
	})

	t.Run("Minimum value", func(t *testing.T) {
		bst := prepareTree()
		require.Equal(t, 20, bst.Min())
	})

	t.Run("Maximum value", func(t *testing.T) {
		bst := prepareTree()
		require.Equal(t, 80, bst.Max())
	})

	t.Run("BST Properties", func(t *testing.T) {
		bst := &BST{}
		values := []int{5, 3, 8, 1, 4, 7, 9}

		for _, v := range values {
			bst.Insert(v)
		}

		rootValue := bst.root.value
		require.Equal(t, 5, rootValue)

		leftChildValue := bst.root.left.value
		require.Equal(t, 3, leftChildValue)

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
