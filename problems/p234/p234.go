package p234

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func isPalindrome(head *ListNode) bool {
	var nodes []*ListNode
	var slow, fast = head, head
	var cnt = 0

	for slow != nil && fast != nil && fast.Next != nil {
		cnt++
		nodes = append(nodes, slow)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
		cnt = cnt*2 + 1
	} else {
		cnt *= 2
	}
	if cnt == 1 {
		return true
	}
	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i].Val != slow.Val {
			return false
		}
		slow = slow.Next
	}
	return true
}
