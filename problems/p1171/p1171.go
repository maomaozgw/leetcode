package p1171

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	sum := 0
	sumMap := make(map[int]*ListNode)
	cur := dummy
	for cur != nil {
		sum += cur.Val
		sumMap[sum] = cur
		cur = cur.Next
	}
	sum = 0
	cur = dummy
	for cur != nil {
		sum += cur.Val
		cur.Next = sumMap[sum].Next
		cur = cur.Next
	}
	return dummy.Next
}
