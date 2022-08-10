package p86

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func partition(head *ListNode, x int) *ListNode {
	var (
		current *ListNode = head
		p1      *ListNode
		p1Head  *ListNode
		p2      *ListNode
		p2Head  *ListNode
	)

	for current != nil {
		if current.Val < x {
			if p1 != nil {
				p1.Next = current
				p1 = current
			} else {
				p1Head = current
				p1 = current
			}
		} else {
			if p2 != nil {
				p2.Next = current
				p2 = current
			} else {
				p2Head = current
				p2 = current
			}
		}
		current = current.Next
	}
	if p1 != nil {
		p1.Next = p2Head
		if p2 != nil {
			p2.Next = nil
		}
		return p1Head
	}
	return p2Head
}
