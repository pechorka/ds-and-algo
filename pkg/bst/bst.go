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
