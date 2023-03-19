package p572

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	}
	if equals(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func equals(one, another *TreeNode) bool {
	if one == nil && another == nil {
		return true
	} else if one == nil || another == nil {
		return false
	}

	if one.Val != another.Val {
		return false
	}
	if !equals(one.Left, another.Left) {
		return false
	}
	if !equals(one.Right, another.Right) {
		return false
	}
	return true
}
