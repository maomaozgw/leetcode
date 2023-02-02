package p437

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	cache := map[int]int{0: 1}

	return dfs(root, cache, root.Val, targetSum)
}

func dfs(root *TreeNode, cache map[int]int, path, target int) int {
	res, diff := 0, path-target

	if cache[diff] > 0 {
		res += cache[diff]
	}

	cache[path]++

	if root.Left != nil {
		res += dfs(root.Left, cache, root.Left.Val+path, target)
	}

	if root.Right != nil {
		res += dfs(root.Right, cache, root.Right.Val+path, target)
	}

	cache[path]--

	return res
}
