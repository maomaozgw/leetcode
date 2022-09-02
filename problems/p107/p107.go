// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

package p107

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func levelOrderBottom(root *TreeNode) [][]int {
	levelMap := map[int][]int{}
	bfs(root, levelMap, 0)
	l := len(levelMap) - 1
	var result = make([][]int, len(levelMap))
	for idx := range levelMap {
		result[l-idx] = levelMap[idx]
	}
	return result
}

func bfs(root *TreeNode, m map[int][]int, l int) {
	if root == nil {
		return

	}
	m[l] = append(m[l], root.Val)
	bfs(root.Left, m, l+1)
	bfs(root.Right, m, l+1)

}
