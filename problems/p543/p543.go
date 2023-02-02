package p543

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result = 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := dfs(root.Left), dfs(root.Right)
		result = max(result, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
