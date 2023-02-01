package p110

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, leftB := getWeight(root.Left)
	if !leftB {
		return false
	}
	right, rightB := getWeight(root.Right)
	if !rightB {
		return false
	}
	return abs(left-right) <= 1
}

func getWeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, leftB := getWeight(root.Left)
	if !leftB {
		return 0, leftB
	}
	right, rightB := getWeight(root.Right)
	if !rightB {
		return 0, rightB
	}
	if abs(left-right) > 1 {
		return 0, false
	}
	return max(left+1, right+1), true
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
