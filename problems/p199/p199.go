package p199

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func rightSideView(root *TreeNode) []int {
	levelMap := map[int]int{}
	dfs(root, levelMap, 0)
	var result = make([]int, len(levelMap))
	for k, v := range levelMap {
		result[k] = v
	}
	return result
}

func dfs(root *TreeNode, m map[int]int, l int) {
	if root == nil {
		return

	}
	m[l] = root.Val
	dfs(root.Left, m, l+1)
	dfs(root.Right, m, l+1)
}
