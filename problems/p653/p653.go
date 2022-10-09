package p653

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func findTarget(root *TreeNode, k int) bool {
	max := findMax(root)
	min := findMin(root)
	if k >= max*2 {
		return false
	}
	if k <= min*2 {
		return false
	}
	return dfsFind(root, root, k)
}

func dfsFind(root, current *TreeNode, k int) bool {
	if current == nil {
		return false
	}

	remain := k - current.Val
	if current.Val != remain && findT(root, remain) {
		return true
	}
	result := dfsFind(root, current.Left, k)
	if result {
		return true
	}
	result = dfsFind(root, current.Right, k)
	if result {
		return true
	}
	return false
}

func findT(root *TreeNode, k int) bool {
	current := root
	for current != nil {
		if current.Val == k {
			return true
		} else if k > current.Val {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return false
}

func findMin(root *TreeNode) int {
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current.Val
}

func findMax(root *TreeNode) int {
	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current.Val
}
