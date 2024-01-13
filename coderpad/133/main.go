package main

import "fmt"

/*
This problem was asked by Amazon.

Given a node in a binary search tree, return the next bigger element, also known as the inorder successor.

For example, the inorder successor of 22 is 30.

	   10
	  /  \
	 5    30
		 /  \
	   22    35

You can assume each node has a parent pointer.
*/
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type InorderSuccessorNode struct {
	Node *Node
}

func searchInorderSuccessor(node *Node) int {
	if node.Right != nil {
		return findMin(node.Right)
	}

	parent := node.Parent
	for parent != nil && node == parent.Right {
		node = parent
		parent = parent.Parent
	}

	return parent.Value
}

func findMin(node *Node) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Value
}

func main() {
	root := &Node{Value: 10, Parent: nil}
	root.Left = &Node{Value: 5, Parent: root}
	root.Right = &Node{Value: 30, Parent: root}
	root.Right.Left = &Node{Value: 22, Parent: root.Right}
	root.Right.Right = &Node{Value: 35, Parent: root.Right}

	// Finding the inorder successor for the node with value 22
	inorderSuccessor := searchInorderSuccessor(root.Right.Left)
	if inorderSuccessor != 0 {
		fmt.Printf("Inorder Successor of 22 is: %d\n", inorderSuccessor)
	} else {
		fmt.Println("No inorder successor found.")
	}
}
