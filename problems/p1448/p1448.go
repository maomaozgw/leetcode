// https://leetcode.com/problems/count-good-nodes-in-binary-tree/

package p1448

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result = 1
	result += dfsNode(root.Left, root.Val)
	result += dfsNode(root.Right, root.Val)
	return result
}

func dfsNode(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}
	var result int
	if root.Val >= maxVal {
		maxVal = root.Val
		result++
	}
	result += dfsNode(root.Left, maxVal)
	result += dfsNode(root.Right, maxVal)
	
	return result
}
