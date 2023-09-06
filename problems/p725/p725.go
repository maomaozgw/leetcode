package p725

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func splitListToParts(head *ListNode, k int) []*ListNode {
	var result = make([]*ListNode, k)
	var current = head
	var cnt = 0
	for current != nil {
		cnt++
		current = current.Next
	}
	if cnt < k {
		current = head
		for i := 0; i < cnt; i++ {
			result[i] = current
			current = current.Next
			result[i].Next = nil
		}
		return result
	}
	var step = cnt / k
	var moreCnt = cnt % k
	current = head
	for i := 0; i < k; i++ {
		var currentStep = step
		if i < moreCnt {
			currentStep++
		}
		result[i] = current
		for j := 0; j < currentStep-1; j++ {
			current = current.Next
		}
		tmp := current
		current = current.Next
		tmp.Next = nil
	}
	return result
}
