// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

package p103

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func zigzagLevelOrder(root *TreeNode) [][]int {
	levelMap := map[int][]int{}
	bfs(root, levelMap, 0)
	var result = make([][]int, len(levelMap))
	for idx := range levelMap {
		if idx%2 == 0 {
			result[idx] = levelMap[idx]
		} else {
			result[idx] = levelMap[idx]
			reverse(result[idx])
		}

	}
	return result
}

func reverse(v []int) {
	l := len(v)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		v[i], v[j] = v[j], v[i]
	}
}

func bfs(root *TreeNode, m map[int][]int, l int) {
	if root == nil {
		return

	}
	m[l] = append(m[l], root.Val)
	bfs(root.Left, m, l+1)
	bfs(root.Right, m, l+1)

}
