package p82

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var lastNum = head.Val
	var lastCnt = 1
	var newH *ListNode
	var newLast *ListNode
	var cNode = head.Next
	for cNode != nil {
		if cNode.Val == lastNum {
			lastCnt++
			cNode = cNode.Next
			continue
		}
		if lastCnt == 1 {
			if newH == nil {
				newH = &ListNode{
					Val:  lastNum,
					Next: nil,
				}
				newLast = newH
			} else {
				newLast.Next = &ListNode{
					Val: lastNum,
				}
				newLast = newLast.Next
			}
		}
		lastCnt = 1
		lastNum = cNode.Val
		cNode = cNode.Next
	}
	if newH == nil && lastCnt == 1 {
		return &ListNode{Val: lastNum}
	}
	if lastCnt == 1 {
		newLast.Next = &ListNode{
			Val: lastNum,
		}
	}
	return newH
}
