package p61

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func rotateRight(head *ListNode, k int) *ListNode {
	var current = head
	var l = 0
	for current != nil {
		current = current.Next
		l++
	}
	if l == 0 {
		return head
	}
	k %= l
	if k == 0 {
		return head
	}
	current = head
	for i := 0; i < l-k-1; i++ {
		current = current.Next
	}
	var newHead = current.Next
	current.Next = nil
	current = newHead
	for current.Next != nil {
		current = current.Next
	}
	current.Next = head
	return newHead
}
