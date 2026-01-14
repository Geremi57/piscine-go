package main

import "fmt"

type Node struct {
	Value int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.left = insert(root.left, value)
	}
	if value > root.Value {
		root.right = insert(root.right, value)
	}
	return root
}

func BTreeApplyInorder(root *Node, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.left, f)
		f(root.Value)
		BTreeApplyInorder(root.right, f)
	}
}

func main() {
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	var root *Node

	for _, val := range values {
		root = insert(root, val)
	}
	
	BTreeApplyInorder(root, func(v int) {
		fmt.Print(v, " ")
	})

}