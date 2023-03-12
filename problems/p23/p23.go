package p23

import (
	"container/heap"

	"github.com/maomaozgw/leetcode/structures/listnode"
)

type ListNode = listnode.G[int]

type Heap []*ListNode

func (h *Heap) Len() int { return len(*h) }

func (h *Heap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	var h = Heap{}
	heap.Init(&h)
	for idx := range lists {
		if lists[idx] == nil {
			continue
		}
		heap.Push(&h, lists[idx])
	}
	if h.Len() == 0 {
		return nil
	}
	item := heap.Pop(&h).(*ListNode)
	var retVal = &ListNode{
		Val: item.Val,
	}
	if item.Next != nil {
		heap.Push(&h, item.Next)
	}
	var current = retVal
	for h.Len() > 0 {
		item = heap.Pop(&h).(*ListNode)
		current.Next = &ListNode{
			Val: item.Val,
		}
		current = current.Next
		if item.Next != nil {
			heap.Push(&h, item.Next)
		}
	}
	return retVal
}
