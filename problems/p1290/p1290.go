package p1290

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func getDecimalValue(head *ListNode) int {
	var result = 0
	for head != nil {
		result <<= 1
		result |= head.Val
		head = head.Next
	}
	return result
}
