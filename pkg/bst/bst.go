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
	} else if value < n.value {
		return n.left.search(value)
	} else {
		return n.right.search(value)
	}
}

func (n *node) delete(value int) *node {
	if n == nil {
		return nil
	}

	if value < n.value {
		n.left = n.left.delete(value)
	} else if value > n.value {
		n.right = n.right.delete(value)
	} else {
		return n.deleteNode()
	}
	return n
}

func (n *node) deleteNode() *node {
	// Zero children
	if n.left == nil && n.right == nil {
		return nil
	}

	// One child
	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	// Two children
	smallestValue := n.right.findSmallestValue()
	n.value = smallestValue
	n.right = n.right.delete(smallestValue)
	return n
}

// findSmallestValue assumes that n is not nil
func (n *node) findSmallestValue() int {
	temp := n
	for temp.left != nil {
		temp = temp.left
	}
	return temp.value
}
