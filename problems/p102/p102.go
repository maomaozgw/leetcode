// https://leetcode.com/problems/binary-tree-level-order-traversal/

package p102

import "github.com/maomaozgw/leetcode/structures/treenode"

type TreeNode = treenode.BinaryTree[int]

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var newQ []*TreeNode
		var subQ []int
		for _, item := range q {
			subQ = append(subQ, item.Val)
			if item.Left != nil {
				newQ = append(newQ, item.Left)

			}
			if item.Right != nil {
				newQ = append(newQ, item.Right)
			}
		}
		result = append(result, subQ)
		q = newQ
	}
	return result
}
