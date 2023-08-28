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
