// https://leetcode.com/problems/n-ary-tree-level-order-traversal/

package p429

import "github.com/maomaozgw/leetcode/structures/treenode"

type Node = treenode.NAtrTree[int]

func levelOrder(root *Node) [][]int {
	m := map[int][]int{}
	if root == nil {
		return nil
	}
	m[0] = append(m[0], root.Val)
	bfs(root, m, 1)
	result := make([][]int, len(m))
	for idx := range m {
		result[idx] = m[idx]
	}
	return result

}

func bfs(root *Node, m map[int][]int, l int) {
	for idx := range root.Children {
		m[l] = append(m[l], root.Children[idx].Val)
	}
	for idx := range root.Children {
		bfs(root.Children[idx], m, l+1)
	}
}
