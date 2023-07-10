package p111

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result = 1
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var newQ []*TreeNode
		for i := 0; i < len(q); i++ {
			if q[i].Left == nil && q[i].Right == nil {
				return result
			}
			if q[i].Left != nil {
				newQ = append(newQ, q[i].Left)
			}
			if q[i].Right != nil {
				newQ = append(newQ, q[i].Right)
			}
		}
		result++
		q = newQ
	}
	return result
}
