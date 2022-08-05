package p623

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := &TreeNode{
			Val:  val,
			Left: root,
		}
		return node
	}
	depth--
	findAndAddByDepth(root, val, depth)
	return root
}

func findAndAddByDepth(root *TreeNode, val int, depth int) {

	if depth == 1 {
		tmpLeft, tmpRight := root.Left, root.Right
		root.Left = &TreeNode{
			Val:  val,
			Left: tmpLeft,
		}
		root.Right = &TreeNode{
			Val:   val,
			Right: tmpRight,
		}
		return
	}
	depth--
	if root.Left != nil {
		findAndAddByDepth(root.Left, val, depth)
	}
	if root.Right != nil {
		findAndAddByDepth(root.Right, val, depth)
	}
}
