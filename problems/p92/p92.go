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
			leftNode = current.N
		}
		if index == right {
			nextAfterRight = current.N
			break
		}
		current = current.N
	}
	newHead, newTail := reverse(leftNode, nextAfterRight)
	if beforeLeft != nil {
		beforeLeft.N = newHead
	} else {
		head = newHead
	}
	newTail.N = nextAfterRight
	return head
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var (
		newHead *ListNode
		current *ListNode = head
		prev    *ListNode
		next    *ListNode
	)
	for current.N != tail {
		next = current.N
		current.N = prev
		prev = current
		current = next
	}
	current.N = prev
	newHead = current

	return newHead, head
}

func main() {
	head := listnode.NewG[int](1, 2, 3, 4, 5)
	listnode.Print(head)

	head = reverseBetween(head, 1, 5)
	listnode.Print(head)
}
