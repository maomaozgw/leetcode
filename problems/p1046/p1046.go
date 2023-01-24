package p1046

import "container/heap"

type Heap []int

func (h *Heap) Len() int { return len(*h) }

func (h *Heap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := Heap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		v1 := heap.Pop(&h).(int)
		v2 := heap.Pop(&h).(int)
		if v1 == v2 {
			continue
		}
		remain := v1 - v2
		heap.Push(&h, remain)
	}
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}
