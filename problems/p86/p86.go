package p86

import "code.byted.org/zhaoguowei/leetcode/structures/listnode"

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
		if current.V < x {
			if p1 != nil {
				p1.N = current
				p1 = current
			} else {
				p1Head = current
				p1 = current
			}
		} else {
			if p2 != nil {
				p2.N = current
				p2 = current
			} else {
				p2Head = current
				p2 = current
			}
		}
		current = current.N
	}
	if p1 != nil {
		p1.N = p2Head
		if p2 != nil {
			p2.N = nil
		}
		return p1Head
	}
	return p2Head
}
