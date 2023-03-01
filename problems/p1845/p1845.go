package p1845

import "container/heap"

type Heap []int

func (h *Heap) Len() int { return len(*h) }

func (h *Heap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
	h *Heap
}

func Constructor(n int) SeatManager {
	var data = make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i + 1
	}
	h := Heap(data)
	heap.Init(&h)
	return SeatManager{
		h: &h,
	}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.h).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.h, seatNumber)
}
