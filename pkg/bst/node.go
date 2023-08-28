package bst

import (
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

func newNode(value int) *node {
	return &node{value: value}
}

func (n *node) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = newNode(value)
		} else {
			n.left.insert(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = newNode(value)
		} else {
			n.right.insert(value)
		}
	}
	// Ignore duplicate values
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
	smallestValue := n.right.min()
	n.value = smallestValue
	n.right = n.right.delete(smallestValue)
	return n
}

// example output:
// 50
// ├── 30
// │   ├── 20
// │   └── 40
// └── 70
//     ├── 60
//     └── 80

func (n *node) draw() string {
	builder := &strings.Builder{}
	n.drawNode("", "", builder)
	return builder.String()
}

func (n *node) drawNode(prefix, childPrefix string, builder *strings.Builder) {
	if n == nil {
		return
	}

	builder.WriteString(prefix)
	builder.WriteString(strconv.Itoa(n.value))
	builder.WriteString("\n")

	n.left.drawNode(childPrefix+"├── ", childPrefix+"│   ", builder)
	n.right.drawNode(childPrefix+"└── ", childPrefix+"    ", builder)
}

func (n *node) traverseInOrder() []int {
	var values []int
	if n.left != nil {
		values = append(values, n.left.traverseInOrder()...)
	}
	values = append(values, n.value)
	if n.right != nil {
		values = append(values, n.right.traverseInOrder()...)
	}
	return values
}

func (n *node) traversePreOrder() []int {
	var values []int
	values = append(values, n.value)
	if n.left != nil {
		values = append(values, n.left.traversePreOrder()...)
	}
	if n.right != nil {
		values = append(values, n.right.traversePreOrder()...)
	}
	return values
}

func (n *node) traversePostOrder() []int {
	var values []int
	if n.left != nil {
		values = append(values, n.left.traversePostOrder()...)
	}
	if n.right != nil {
		values = append(values, n.right.traversePostOrder()...)
	}
	values = append(values, n.value)
	return values
}

func (n *node) traverseLevelOrder() []int {
	var values []int
	var queue []*node

	queue = append(queue, n)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		values = append(values, current.value)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return values
}

func (n *node) height() int {
	if n == nil {
		return 0
	}

	leftHeight := n.left.height()
	rightHeight := n.right.height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (n *node) max() int {
	temp := n
	for temp.right != nil {
		temp = temp.right
	}
	return temp.value
}

func (n *node) min() int {
	temp := n
	for temp.left != nil {
		temp = temp.left
	}
	return temp.value
}
