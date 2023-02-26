package p404

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum = sumLeft(root.Left)
	sum += sumRight(root.Right)
	return sum
}

func sumLeft(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return sumOfLeftLeaves(root)
}

func sumRight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	return sumOfLeftLeaves(root)
}
