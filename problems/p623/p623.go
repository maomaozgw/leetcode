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
	targetNodes := findNodesByDepth(root, depth)
	for _, node := range targetNodes {
		if node == nil {
			continue
		}
		tmpLeft, tmpRight := node.Left, node.Right
		node.Left = &TreeNode{
			Val:  val,
			Left: tmpLeft,
		}
		node.Right = &TreeNode{
			Val:   val,
			Right: tmpRight,
		}

	}
	return root
}

func findNodesByDepth(root *TreeNode, depth int) []*TreeNode {
	depth--
	if depth == 0 {
		return []*TreeNode{root}
	}
	var result []*TreeNode
	if root.Left != nil {
		result = append(result, findNodesByDepth(root.Left, depth)...)
	}
	if root.Right != nil {
		result = append(result, findNodesByDepth(root.Right, depth)...)
	}
	return result
}
