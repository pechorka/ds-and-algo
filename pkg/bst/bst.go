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
func (n *BST) TraverseInOrder() []int {
	return n.root.traverseInOrder()
}

// TraversePreOrder root, left, right (DFS)
func (n *BST) TraversePreOrder() []int {
	return n.root.traversePreOrder()
}

// TraversePostOrder left, right, root (DFS)
func (bst *BST) TraversePostOrder() []int {
	return bst.root.traversePostOrder()
}

// TraverseLevelOrder level by level (BFS)
func (bst *BST) TraverseLevelOrder() []int {
	return bst.root.traverseLevelOrder()
}
