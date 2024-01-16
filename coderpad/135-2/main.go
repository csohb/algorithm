package main

import (
	"fmt"
	"math"
)

/*
   트리에서 최소값, 최대값 찾는 function
*/

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type MinNode struct {
	Node     *Node
	MinValue int
}

type MaxNode struct {
	Node     *Node
	MaxValue int
}

func findMinNode(root *Node) int {
	// bfs
	var queue []*MinNode
	queue = append(queue, &MinNode{
		Node:     root,
		MinValue: root.Value,
	})

	min := math.MaxInt64

	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		node := curNode.Node

		if node.Left != nil {
			queue = append(queue, &MinNode{
				Node:     node.Left,
				MinValue: node.Left.Value,
			})
		}

		if node.Right != nil {
			queue = append(queue, &MinNode{
				Node:     node.Right,
				MinValue: node.Right.Value,
			})
		}

		if node.Value < min {
			min = node.Value
		}
	}
	return min
}

func findMaxNode(root *Node) int {
	// bfs
	var queue []*MaxNode
	queue = append(queue, &MaxNode{
		Node:     root,
		MaxValue: root.Value,
	})

	max := math.MinInt

	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		node := curNode.Node

		if node.Left != nil {
			queue = append(queue, &MaxNode{
				Node:     node.Left,
				MaxValue: node.Left.Value,
			})
		}

		if node.Right != nil {
			queue = append(queue, &MaxNode{
				Node:     node.Right,
				MaxValue: node.Right.Value,
			})
		}

		if node.Value > max {
			max = node.Value
		}
	}
	return max
}

func main() {
	root := &Node{Value: 9, Left: &Node{Value: 7, Left: &Node{Value: -100, Left: nil, Right: nil}, Right: nil},
		Right: &Node{Value: 2,
			Left:  &Node{Value: 6, Left: &Node{Value: 1, Left: nil, Right: nil}, Right: &Node{Value: 14, Left: nil, Right: nil}},
			Right: &Node{Value: 3, Left: &Node{Value: -1, Left: nil, Right: &Node{Value: 8, Left: nil, Right: nil}}, Right: nil}}}

	min := findMinNode(root)
	max := findMaxNode(root)

	fmt.Println("min:", min)
	fmt.Println("max:", max)
}
