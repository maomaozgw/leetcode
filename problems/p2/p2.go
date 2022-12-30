package p2

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = l1
	var flag = 0
	var lastL1 = head
	for l1 != nil && l2 != nil {
		l1.Val += l2.Val + flag
		flag = l1.Val / 10
		l1.Val %= 10
		lastL1 = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	for l2 != nil {
		lastL1.Next = &ListNode{
			Val: l2.Val + flag,
		}
		l2 = l2.Next
		lastL1 = lastL1.Next
		flag = lastL1.Val / 10
		lastL1.Val %= 10
	}
	for lastL1.Next != nil {
		lastL1.Next.Val += flag
		lastL1 = lastL1.Next
		flag = lastL1.Val / 10
		lastL1.Val %= 10
	}
	if flag > 0 {
		lastL1.Next = &ListNode{
			Val: flag,
		}
	}

	return head
}
