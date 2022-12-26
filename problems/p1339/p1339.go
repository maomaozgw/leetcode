package p1339

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

const (
	modVal = 1000000007
)

func maxProduct(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sumMap = map[*TreeNode]int{}
	var total = int64(dfs(root, sumMap))
	if len(sumMap) == 1 {
		// only root
		return int(total % modVal)
	}
	var max = int64(0)
	delete(sumMap, root)
	for _, val := range sumMap {
		productVal := (total - int64(val)) * int64(val)
		if productVal > max {
			max = productVal
		}
	}
	return int(max % modVal)
}

func dfs(root *TreeNode, sumMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	val := root.Val
	val += dfs(root.Left, sumMap)
	val += dfs(root.Right, sumMap)
	sumMap[root] = val
	return val
}
