package p515

import (
	"math"

	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	var q []*TreeNode

	q = append(q, root)

	for len(q) > 0 {
		levelMax := math.MinInt
		var newQ []*TreeNode
		for _, item := range q {
			levelMax = max(levelMax, item.Val)
			if item.Left != nil {
				newQ = append(newQ, item.Left)
			}
			if item.Right != nil {
				newQ = append(newQ, item.Right)
			}
		}
		q = newQ
		result = append(result, levelMax)
	}

	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
