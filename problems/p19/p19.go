// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package p19

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	var current = head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}
	if len(nodes) <= 1 {
		return nil
	}
	if n == len(nodes) {
		return head.Next
	}
	idx := len(nodes) - n - 1
	nodes[idx].Next = nodes[idx+1].Next
	return head
}
