// https://leetcode.com/problems/validate-binary-search-tree/

package p98

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
	"math"
)

type TreeNode = treenode.BinaryTree[int]

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt, math.MaxInt)
}

func valid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}
	return valid(root.Left, min, root.Val) && valid(root.Right, root.Val, max)
}

func walk(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	nums = append(nums, walk(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, walk(root.Right)...)
	return nums
}
