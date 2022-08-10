package p108

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	} else if len(nums) == 2 {
		return &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val: nums[0],
			},
		}
	} else if len(nums) == 3 {
		return &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val: nums[0],
			},
			Right: &TreeNode{
				Val: nums[2],
			},
		}
	}
	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
