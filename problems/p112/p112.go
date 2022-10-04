// https://leetcode.com/problems/path-sum/

package p112

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, targetSum)
}

func dfs(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}

	val -= root.Val

	if root.Left == nil && root.Right == nil {
		return val == 0
	}

	result := dfs(root.Left, val)
	if result {
		return result
	}
	result = dfs(root.Right, val)
	if result {
		return result
	}
	return false
}
