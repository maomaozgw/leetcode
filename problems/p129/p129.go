package p129

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func sumNumbers(root *TreeNode) int {
	var sum = 0
	var sumPath func(*TreeNode, int)
	sumPath = func(tn *TreeNode, i int) {
		if tn.Left == nil && tn.Right == nil {
			i = i*10 + tn.Val
			sum += i
			return
		}
		i = i*10 + tn.Val
		if tn.Left != nil {
			sumPath(tn.Left, i)
		}
		if tn.Right != nil {
			sumPath(tn.Right, i)
		}
	}
	sumPath(root, 0)
	return sum
}
