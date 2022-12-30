package p142

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	/*
			 1 -> 2 -> 3 -> 4
			      ⬆️ <- <- ⬇️
		    1->2 =K
		    2->4->2 = M
			slow: X  --> X = n*M
			fast: 2X = X + n*M = 2*n*M
			s2: walk  K, slow = K + n*M 意味着两个指针会在 K 处重合，因为 K + n*M 必然在 K 处

	*/
	var fast = head
	var slow = head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// has cycle
			var s2 = head
			for s2 != slow {
				slow = slow.Next
				s2 = s2.Next
			}
			return slow
		}
	}
	//no cycle
	return nil

}
