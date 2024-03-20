package p1669

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var head = list2
	var cur *ListNode
	var end *ListNode
	if a > 0 {
		head = list1
		cur = list1
		for i := 1; i < a; i++ {
			cur = cur.Next
		}
		end = cur.Next
		cur.Next = list2
		cur = cur.Next
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	for i := 0; i < b-a; i++ {
		end = end.Next
	}
	cur.Next = end.Next
	return head
}
