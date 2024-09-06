package p3217

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func modifiedList(nums []int, head *ListNode) *ListNode {
	var excludeMap = make(map[int]bool)
	for _, num := range nums {
		excludeMap[num] = true
	}
	var h = &ListNode{Next: nil}
	var c = h
	for head != nil {
		if excludeMap[head.Val] {
			head = head.Next
			continue
		}
		c.Next = head
		head = head.Next
		c = c.Next
		c.Next = nil
	}
	return h.Next
}
