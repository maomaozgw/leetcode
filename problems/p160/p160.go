package p160

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	/*
		la : a + m
		lb : b + m
	*/
	var ca, cb = headA, headB
	for ca != nil && cb != nil {
		ca = ca.Next
		cb = cb.Next
	}
	var cMore = ca
	var hMore = headA
	var hLess = headB
	if ca == nil {
		cMore = cb
		hMore = headB
		hLess = headA
	}
	for cMore != nil {
		// walked b+m, remain b-a

		hMore = hMore.Next
		cMore = cMore.Next
	}
	for hMore != nil && hLess != nil && hMore != hLess {
		hMore = hMore.Next
		hLess = hLess.Next
	}
	if hMore == nil || hLess == nil {
		return nil
	}
	return hMore
}
