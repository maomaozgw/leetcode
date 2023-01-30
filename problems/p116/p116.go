package p116

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		var newQ []*Node
		var i int
		for ; i < len(q)-1; i++ {
			if q[i].Left != nil {
				newQ = append(newQ, q[i].Left)
			}
			if q[i].Right != nil {
				newQ = append(newQ, q[i].Right)
			}
			q[i].Next = q[i+1]
		}
		if q[i].Left != nil {
			newQ = append(newQ, q[i].Left)
		}
		if q[i].Right != nil {
			newQ = append(newQ, q[i].Right)
		}
		q = newQ
	}
	return root
}
