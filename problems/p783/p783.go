// https://leetcode.com/problems/minimum-distance-between-bst-nodes/

package p783

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func minDiffInBST(root *TreeNode) int {
	var values []int
	dfs(root, &values)
	var minGap = 1000000
	for i := 1; i < len(values); i++ {
		gap := values[i] - values[i-1]
		if gap < minGap {
			minGap = gap
		}
	}
	return minGap
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, result)
	*result = append(*result, root.Val)
	dfs(root.Right, result)

}
