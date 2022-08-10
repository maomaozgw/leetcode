// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/

package p109

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type TreeNode = treenode.BinaryTree[int]
type ListNode = listnode.G[int]

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func buildTree(start, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}
	var oneStep = start
	var twoStep = start

	for oneStep != end && twoStep != end && twoStep.Next != end {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
	}
	root := &TreeNode{
		Val: oneStep.Val,
	}
	// start -> Middle
	root.Left = buildTree(start, oneStep)
	// Middle -> End
	root.Right = buildTree(oneStep.Next, end)
	return root
}
