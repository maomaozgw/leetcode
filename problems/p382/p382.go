package p382

import (
	"math/rand"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

type Solution struct {
	lists []*ListNode
}

func Constructor(head *ListNode) Solution {
	s := Solution{
		lists: []*ListNode{},
	}
	for head != nil {
		s.lists = append(s.lists, head)
		head = head.Next
	}
	return s
}

func (this *Solution) GetRandom() int {
	idx := rand.Intn(len(this.lists))
	return this.lists[idx].Val
}
