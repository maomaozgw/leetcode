// https://leetcode.com/problems/average-of-levels-in-binary-tree/

package p637

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func averageOfLevels(root *TreeNode) []float64 {
	sum := map[int]float64{}
	count := map[int]int{}
	sum[0] = float64(root.Val)
	count[0] = 1
	bfs(root, sum, count, 1)
	var result = make([]float64, len(count))
	for idx := range count {
		result[idx] = sum[idx] / float64(count[idx])
	}
	return result
}

func bfs(root *TreeNode, sum map[int]float64, count map[int]int, level int) {

	if root.Left != nil {
		sum[level] += float64(root.Left.Val)
		count[level]++
		bfs(root.Left, sum, count, level+1)
	}
	if root.Right != nil {
		sum[level] += float64(root.Right.Val)
		count[level]++
		bfs(root.Right, sum, count, level+1)
	}
}
