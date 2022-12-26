package p701

import "github.com/maomaozgw/leetcode/structures/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = treenode.BinaryTree[int]

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
			return root
		}
		insertIntoBST(root.Right, val)
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
			return root
		}
		insertIntoBST(root.Left, val)
	}
	return root
}
