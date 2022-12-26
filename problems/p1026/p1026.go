package p1026

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"math"
)

type TreeNode = treenode.BinaryTree[int]

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, math.MinInt, math.MaxInt)
}
func dfs(root *TreeNode, maxVal, minVal int) int {
	if root == nil {
		return 0
	}
	maxVal = max(root.Val, maxVal)
	minVal = min(root.Val, minVal)
	if root.Left == nil && root.Right == nil {
		return maxVal - minVal
	}
	return max(dfs(root.Left, maxVal, minVal), dfs(root.Right, maxVal, minVal))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
