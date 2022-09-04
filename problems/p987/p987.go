// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/

package p987

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"sort"
)

type TreeNode = treenode.BinaryTree[int]

const (
	offset = 10000
)

func verticalTraversal(root *TreeNode) [][]int {
	m := map[int][]int{}
	dfs(root, m, 0, 0)
	result := make([][]int, len(m))
	minColumn := 0
	for column := range m {
		if column < minColumn {
			minColumn = column
		}
	}
	for column := range m {
		result[column-minColumn] = m[column]
		sorts(result[column-minColumn])

	}
	return result
}

func sorts(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		x := nums[i] % offset
		y := nums[j] % offset
		xr := nums[i] / offset
		xy := nums[j] / offset
		if xr == xy {
			return x < y
		}
		return xr < xy
	})
	for i := range nums {
		nums[i] = nums[i] % offset
	}
}

func dfs(root *TreeNode, m map[int][]int, column, row int) {
	if root == nil {
		return
	}
	m[column] = append(m[column], row*offset+root.Val)
	dfs(root.Left, m, column-1, row+1)
	dfs(root.Right, m, column+1, row+1)
}
