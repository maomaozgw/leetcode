package p2807

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	node := head
	for node.Next != nil {
		node.Next = &ListNode{
			Val:  gcd(node.Val, node.Next.Val),
			Next: node.Next,
		}
		node = node.Next.Next
	}
	return head
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
