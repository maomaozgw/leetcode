package p203

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func removeElements(head *ListNode, val int) *ListNode {
	var root, prev, current *ListNode
	current = head
	for current != nil {
		if current.Val != val {
			root = current
			prev = current
			current = current.Next
			break
		}
		current = current.Next
	}
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
			current = current.Next
			continue
		}
		prev = current
		current = current.Next
	}
	return root
}
