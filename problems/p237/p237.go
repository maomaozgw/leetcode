package p237

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	next.Next = nil
}
