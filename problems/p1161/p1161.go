package p1161

import (
	"math"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q []*TreeNode
	q = append(q, root)
	var result = 0
	var maxSum = math.MinInt
	var level = 1
	for len(q) > 0 {
		var newQ []*TreeNode
		var levelSum = 0
		for i := 0; i < len(q); i++ {
			c := q[i]
			levelSum += c.Val
			if c.Left != nil {
				newQ = append(newQ, c.Left)
			}
			if c.Right != nil {
				newQ = append(newQ, c.Right)
			}
		}
		if levelSum > maxSum {
			result = level
			maxSum = levelSum
		}
		q = newQ
		level++
	}
	return result
}
