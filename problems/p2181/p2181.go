package p2181

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func mergeNodes(head *ListNode) *ListNode {
	var left = head
	var right = head.Next
	for right.Next != nil {
		if right.Val != 0 {
			left.Val += right.Val
			right = right.Next
			continue
		}
		left = left.Next
		left.Val = 0
		right = right.Next
	}
	left.Next = nil
	return head
}
