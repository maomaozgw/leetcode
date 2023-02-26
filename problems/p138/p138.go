package p138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var nodeMap = map[*Node]*Node{}
	var current = head
	var newHead = &Node{
		Val: head.Val,
	}
	var newCur = newHead
	nodeMap[head] = newHead
	current = current.Next
	for current != nil {
		newVal := &Node{
			Val: current.Val,
		}
		newCur.Next = newVal
		nodeMap[current] = newVal
		current = current.Next
		newCur = newCur.Next
	}
	newCur = newHead
	current = head
	for current != nil {
		newCur.Random = nodeMap[current.Random]
		current = current.Next
		newCur = newCur.Next
	}
	return newHead
}
