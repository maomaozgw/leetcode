package p2415

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var q []*TreeNode
	q = append(q, root)
	level := 0
	for len(q) > 0 {
		newQ := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			c := q[i]
			if c.Left != nil {
				newQ = append(newQ, c.Left)
			}
			if c.Right != nil {
				newQ = append(newQ, c.Right)
			}
		}
		q = newQ
		level++
		if level%2 == 1 {
			for i := 0; i < len(q)/2; i++ {
				q[i].Val, q[len(q)-1-i].Val = q[len(q)-1-i].Val, q[i].Val
			}
		}
	}
	return root
}
