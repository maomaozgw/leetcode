package p1372

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func longestZigZag(root *TreeNode) int {
	var dfs func(node *TreeNode) (leftLen, rightLen, maxLen int)
	dfs = func(node *TreeNode) (leftLen int, rightLen int, maxLen int) {
		if node == nil {
			return -1, -1, -1
		}

		var _, lrl, lml = dfs(node.Left)
		var rll, _, rml = dfs(node.Right)

		lrl++
		rll++

		var maxL = max(lrl, rll, lml, rml)
		return lrl, rll, maxL
	}
	_, _, result := dfs(root)
	return result
}

func max(nums ...int) int {
	var maxN = 0
	for i := range nums {
		if nums[i] > maxN {
			maxN = nums[i]
		}
	}
	return maxN
}
