package main

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
	children []*Node
	value    int
}

// BFS Breadth-first search <->
func findById(root *Node, searchValue int) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.value == searchValue {
			return nextUp
		}
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

func findByIdDFS(node *Node, searchValue int) *Node {
	if node.value == searchValue {
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			findByIdDFS(child, searchValue)
		}
	}
	return nil
}

func main() {

}
