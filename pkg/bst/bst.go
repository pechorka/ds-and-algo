package bst

type BST struct {
	root *node
}

func New() *BST {
	return &BST{}
}

func (bst *BST) Insert(value int) {
	if bst.root == nil {
		bst.root = newNode(value)
		return
	}

	bst.root.insert(value)
}

func (bst *BST) Search(value int) bool {
	return bst.root.search(value)
}

func (bst *BST) Delete(value int) {
	bst.root = bst.root.delete(value)
}

func (bst *BST) DrawTree() string {
	return bst.root.draw()
}

// TraverseInOrder left, root, right (DFS)
func (bst *BST) TraverseInOrder() []int {
	return bst.root.traverseInOrder()
}

// TraversePreOrder root, left, right (DFS)
func (bst *BST) TraversePreOrder() []int {
	return bst.root.traversePreOrder()
}

// TraversePostOrder left, right, root (DFS)
func (bst *BST) TraversePostOrder() []int {
	return bst.root.traversePostOrder()
}

// TraverseLevelOrder level by level (BFS)
func (bst *BST) TraverseLevelOrder() []int {
	return bst.root.traverseLevelOrder()
}

func (bst *BST) Height() int {
	return bst.root.height()
}

func (bst *BST) Max() int {
	return bst.root.max()
}

func (bst *BST) Min() int {
	return bst.root.min()
}
