package p1962

import (
	"container/heap"
)

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

func minStoneSum(piles []int, k int) int {
	h := Heap(piles)
	heap.Init(&h)
	for i := 0; i < k; i++ {
		val := heap.Pop(&h).(int)
		if val%2 == 0 {
			val /= 2
		} else {
			val = val/2 + 1
		}
		heap.Push(&h, val)
	}
	var sum = 0
	for i := 0; i < len(piles); i++ {
		sum += piles[i]
	}
	return sum
}
