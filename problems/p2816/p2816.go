package p2816

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func doubleIt(head *ListNode) *ListNode {
	carry := doubleNode(head)
	if carry > 0 {
		return &ListNode{
			Val:  carry,
			Next: head,
		}
	}
	return head
}

func doubleNode(node *ListNode) int {
	if node.Next == nil {
		node.Val *= 2
		if node.Val >= 10 {
			node.Val -= 10
			return 1
		}
		return 0
	}
	carry := doubleNode(node.Next)
	node.Val *= 2
	node.Val += carry
	if node.Val >= 10 {
		node.Val -= 10
		return 1
	}
	return 0
}
