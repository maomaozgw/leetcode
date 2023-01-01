package p25

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var nodes []*ListNode
	var c = head
	for i := 0; i < k; i++ {
		nodes = append(nodes, c)
		c = c.Next
		if c == nil {
			break
		}
	}
	if len(nodes) < k {
		return nodes[0]
	}
	nodes[0].Next = nodes[k-1].Next
	for i := 1; i < k; i++ {
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = reverseKGroup(nodes[0].Next, k)
	return nodes[k-1]
}
