package p328

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var evenHead, evenCurrent = head.Next, head.Next
	var oddCurrent = head
	var oddLast, evenLast = oddCurrent, evenCurrent

	for evenCurrent != nil && oddCurrent != nil && evenCurrent.Next != nil && oddCurrent.Next != nil {
		evenCurrent = evenCurrent.Next.Next
		oddCurrent = oddCurrent.Next.Next
		oddLast.Next = oddCurrent
		evenLast.Next = evenCurrent
		oddLast = oddCurrent
		evenLast = evenCurrent
	}
	oddLast.Next = evenHead
	return head
}
