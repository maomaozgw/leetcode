package p1038

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func bstToGst(root *TreeNode) *TreeNode {

	var dfs func(node *TreeNode)
	var preSum int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		node.Val += preSum
		preSum = node.Val
		dfs(node.Left)
	}
	dfs(root)
	return root
}
