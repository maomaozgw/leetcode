package p230

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func kthSmallest(root *TreeNode, k int) int {
	return inorder(root, &k)
}

func inorder(root *TreeNode, k *int) int {
	if root == nil {
		return -1
	}
	if l := inorder(root.Left, k); l != -1 {
		return l
	}
	if *k--; *k == 0 {
		return root.Val
	}
	if r := inorder(root.Right, k); r != -1 {
		return r
	}
	return -1
}

func morrisTraversal(node *TreeNode) *TreeNode {
	for node != nil {
		if node.Left != nil {
			pred := node.Left
			for ; pred.Right != nil && pred.Right != node; pred = pred.Right {
			}
			if pred.Right == nil {
				pred.Right, node = node, node.Left
			} else {
				pred.Right = nil
				return node
			}
		} else {
			return node
		}
	}
	return node
}
