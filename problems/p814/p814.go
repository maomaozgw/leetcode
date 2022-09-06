package p814

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func pruneTree(root *TreeNode) *TreeNode {
	result := dfs(root)
	if !result {
		return nil
	}
	return root
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return false
	}
	leftResult := dfs(root.Left)
	if !leftResult {
		root.Left = nil
	}
	rightResult := dfs(root.Right)
	if !rightResult {
		root.Right = nil
	}
	return leftResult || rightResult || root.Val == 1
}
