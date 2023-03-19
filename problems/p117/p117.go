package p117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		var newQ []*Node
		for i := 0; i < len(q)-1; i++ {
			q[i].Next = q[i+1]
			if q[i].Left != nil {
				newQ = append(newQ, q[i].Left)
			}
			if q[i].Right != nil {
				newQ = append(newQ, q[i].Right)
			}
		}
		var n = q[len(q)-1]
		if n.Left != nil {
			newQ = append(newQ, n.Left)
		}
		if n.Right != nil {
			newQ = append(newQ, n.Right)
		}
		q = newQ
	}
	return root
}
