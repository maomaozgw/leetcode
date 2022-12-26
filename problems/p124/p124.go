package p124

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

const (
	intMin = -1 << 31
)

func maxPathSum(root *TreeNode) int {
	maxSum, _ := dfs(root)
	return maxSum
}

func dfs(root *TreeNode) (maxSum, pathSum int) {
	if root == nil {
		return intMin, intMin
	}
	leftMax, leftPath := dfs(root.Left)
	rightMax, rightPath := dfs(root.Right)
	maxSum = max(max(leftMax, rightMax), root.Val+leftPath+rightPath)
	pathSum = max(root.Val+max(leftPath, rightPath), root.Val)
	maxSum = max(pathSum, maxSum)
	return
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
