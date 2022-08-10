package p92

import (
	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if right <= left {
		return head
	}
	var (
		current        *ListNode = head
		beforeLeft     *ListNode
		leftNode       *ListNode = head
		nextAfterRight *ListNode
		index          = 1
	)
	for ; index <= right; index++ {
		if index+1 == left {
			beforeLeft = current
			leftNode = current.Next
		}
		if index == right {
			nextAfterRight = current.Next
			break
		}
		current = current.Next
	}
	newHead, newTail := reverse(leftNode, nextAfterRight)
	if beforeLeft != nil {
		beforeLeft.Next = newHead
	} else {
		head = newHead
	}
	newTail.Next = nextAfterRight
	return head
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var (
		newHead *ListNode
		current *ListNode = head
		prev    *ListNode
		next    *ListNode
	)
	for current.Next != tail {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	current.Next = prev
	newHead = current

	return newHead, head
}

func main() {
	head := listnode.NewG[int](1, 2, 3, 4, 5)
	listnode.Print(head)

	head = reverseBetween(head, 1, 5)
	listnode.Print(head)
}
