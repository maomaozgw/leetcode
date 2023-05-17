package p2130

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func pairSum(head *ListNode) int {
	var slow, fast = head, head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow
	left := reverse(head, slow)
	var result = 0
	for right != nil && left != nil {
		result = max(result, right.Val+left.Val)
		left = left.Next
		right = right.Next
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func reverse(head *ListNode, end *ListNode) *ListNode {
	var last *ListNode
	var current = head
	for current != end {
		tmp := current.Next
		current.Next = last
		last = current
		current = tmp
	}
	return last
}
