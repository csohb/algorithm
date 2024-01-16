package main

import "fmt"

// TreeNode represents a node in the binary search tree
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
	Parent      *TreeNode
}

// InorderSuccessor finds the inorder successor of a given node in a BST
func InorderSuccessor(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	// If the node has a right child, find the leftmost node in the right subtree
	if node.Right != nil {
		return findLeftmost(node.Right)
	}

	// If the node doesn't have a right child, go up the tree until finding a node
	// where the current node is in the left subtree
	current := node
	parent := node.Parent
	for parent != nil && current == parent.Right {
		current = parent
		parent = parent.Parent
	}

	return parent
}

// findLeftmost finds the leftmost node in the given subtree
func findLeftmost(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {
	// Example usage
	root := &TreeNode{Value: 10}
	root.Left = &TreeNode{Value: 5, Parent: root}
	root.Right = &TreeNode{Value: 30, Parent: root}
	root.Right.Left = &TreeNode{Value: 22, Parent: root.Right}
	root.Right.Right = &TreeNode{Value: 35, Parent: root.Right}
	root.Right.Right.Left = &TreeNode{Value: 27, Parent: root.Right.Right}
	root.Right.Right.Right = &TreeNode{Value: 89, Parent: root.Right.Right}

	node := root.Right.Right.Left // Node with value 27
	successor := InorderSuccessor(node)

	if successor != nil {
		fmt.Printf("Inorder successor of %d is %d\n", node.Value, successor.Value)
	} else {
		fmt.Printf("%d doesn't have an inorder successor\n", node.Value)
	}
}
