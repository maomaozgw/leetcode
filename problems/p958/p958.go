package p958

import (
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]

func isCompleteTree(root *TreeNode) bool {
	var q = []*TreeNode{root}
	var hasNil = false
	for len(q) > 0 {
		var newQ []*TreeNode
		for _, item := range q {
			if item == nil {
				hasNil = true
				continue
			}
			if hasNil {
				return false
			}
			newQ = append(newQ, item.Left)
			newQ = append(newQ, item.Right)
		}
		q = newQ
	}
	return true
}
