package p450

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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	}
	gatherNodes := root.Right
	for gatherNodes.Left != nil {
		gatherNodes = gatherNodes.Left
	}
	gatherNodes.Left = root.Left
	return root.Right
}
