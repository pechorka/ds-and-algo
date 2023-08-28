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
	if bst.root == nil {
		return false
	}

	return bst.root.search(value)
}

type node struct {
	value int
	left  *node
	right *node
}

func newNode(value int) *node {
	return &node{value: value}
}

func (n *node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = newNode(value)
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = newNode(value)
		} else {
			n.right.insert(value)
		}
	}
}

func (n *node) search(value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	if value <= n.value {
		return n.left.search(value)
	}

	return n.right.search(value)
}
