package p1367

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
	"github.com/maomaozgw/leetcode/structures/treenode"
)

type ListNode = listnode.G[int]

type TreeNode = treenode.BinaryTree[int]

func isSubPath(head *ListNode, root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == nil {
			continue
		}
		queue = append(queue, curr.Left)
		queue = append(queue, curr.Right)

		if dfs(head, curr) {
			return true
		}
	}
	return false
}

func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		if head.Next == nil {
			return true
		}
		return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
	}
	return false
}
