package p148

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var slow, fast = head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var middle = slow
	slow = slow.Next
	middle.Next = nil
	var left = sortList(head)
	var right = sortList(slow)
	head = sortedMerge(left, right)
	return head
}

func sortedMerge(left, right *ListNode) *ListNode {
	var result *ListNode
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val <= right.Val {
		result = left
		result.Next = sortedMerge(left.Next, right)
	} else {
		result = right
		result.Next = sortedMerge(left, right.Next)
	}
	return result
}
