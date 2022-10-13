// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

package p2095

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func deleteMiddle(head *ListNode) *ListNode {
	var slow, fast = head, head
	var prev = head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next

		fast = fast.Next.Next
	}
	if prev == slow {
		return nil
	}
	prev.Next = slow.Next
	slow.Next = nil
	return head
}
