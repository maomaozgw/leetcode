package p143

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func reorderList(head *ListNode) {
	var nodes []*ListNode
	var c = head
	for c != nil {
		nodes = append(nodes, c)
		c = c.Next
	}
	if len(nodes) < 3 {
		return
	}
	c = head
	var l = len(nodes) - 1
	for i := 0; i < len(nodes)/2; i++ {
		nodes[i].Next = nodes[l-i]
		nodes[l-i].Next = nodes[i+1]
	}
	if l%2 == 0 {
		nodes[l/2].Next = nil
	} else {
		nodes[l/2+1].Next = nil
	}
}
