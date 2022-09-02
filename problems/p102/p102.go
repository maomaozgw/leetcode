// https://leetcode.com/problems/binary-tree-level-order-traversal/

package p102

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func levelOrder(root *TreeNode) [][]int {
	levelMap := map[int][]int{}
	bfs(root, levelMap, 0)
	var result = make([][]int, len(levelMap))
	for idx := range levelMap {
		result[idx] = levelMap[idx]
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
