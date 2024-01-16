package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type NodeWithSum struct {
	Node *Node
	Sum  int
}

func (n *Node) String() string {
	return fmt.Sprintf("<Node value=%d, left=%v, right=%v>", n.Value, n.Left, n.Right)
}

// 함수는 주어진 이진 트리의 최소 경로 합을 계산
func minPathSum(root *Node) int {
	// BFS를 위한 큐를 초기화하고 루트 노드를 추가한다.
	var queue []*NodeWithSum
	queue = append(queue, &NodeWithSum{
		Node: root,
		Sum:  root.Value,
	})

	fmt.Println("queue first:", queue)

	// 최소값을 초기화하기 위한 변수
	min := math.MaxInt64

	// 큐가 비어질때까지 반복
	for len(queue) > 0 {

		// 큐의 첫번째 노드, 즉 처음 들어온 값(노드)
		nodeWithSum := queue[0]

		// 큐에서 처음 들어온 노드를 제거하기
		queue = queue[1:]

		fmt.Println("queue[0]:", nodeWithSum)
		fmt.Println("queue[1:]:", queue)

		// 현재 노드와 누적된 합을 변수에 저장
		node := nodeWithSum.Node
		currSum := nodeWithSum.Sum

		// 리프 노드에 도달했을 때, 현재까지의 합을 최솟값과 비교하여 업데이트
		if node.Left == nil && node.Right == nil {
			if currSum < min {
				min = currSum
			}
			continue
		}

		// 왼쪽 자식 노드가 존재하면 큐에 추가
		if node.Left != nil {
			queue = append(queue, &NodeWithSum{
				Node: node.Left,
				Sum:  currSum + node.Left.Value,
			})
		}

		// 오른쪽 자식 노드가 존재하면 큐에 추가
		if node.Right != nil {
			queue = append(queue, &NodeWithSum{
				Node: node.Right,
				Sum:  currSum + node.Right.Value,
			})
		}
	}

	return min
}

func main() {
	root := &Node{Value: 10, Left: &Node{Value: 5, Left: nil, Right: &Node{Value: 2, Left: nil, Right: nil}},
		Right: &Node{Value: 5, Left: nil, Right: &Node{Value: 1, Left: &Node{Value: -1, Left: nil, Right: nil}, Right: nil}}}
	fmt.Println(minPathSum(root))
}
