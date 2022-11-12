// https://leetcode.com/problems/find-median-from-data-stream/

package p295

import "container/heap"

type Heap struct {
	data     []int
	lessFunc func(i, j int) bool
}

func (h *Heap) Len() int { return len(h.data) }

func (h *Heap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *Heap) Push(x interface{}) { h.data = append(h.data, x.(int)) }

func (h *Heap) Less(i, j int) bool {
	return h.lessFunc(h.data[i], h.data[j])
}

func (h *Heap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func gather(i, j int) bool {
	return !less(i, j)
}

func less(i, j int) bool {
	if i < j {
		return true
	}
	return false
}

func NewHeap(lessFunc func(i, j int) bool) *Heap {
	return &Heap{
		data:     []int{},
		lessFunc: lessFunc,
	}
}

type MedianFinder struct {
	minHeap *Heap
	maxHeap *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: NewHeap(gather),
		maxHeap: NewHeap(less),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() == 0 || (this.minHeap.Len() != 0 && num <= this.minHeap.data[0]) {
		heap.Push(this.minHeap, num)
		if this.minHeap.Len()-this.maxHeap.Len() > 1 {
			heap.Push(this.maxHeap, heap.Pop(this.minHeap))
		}
		return
	}
	heap.Push(this.maxHeap, num)
	if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	switch {
	case this.minHeap.Len() > this.maxHeap.Len():
		return float64(this.minHeap.data[0])
	case this.minHeap.Len() < this.maxHeap.Len():
		return float64(this.maxHeap.data[0])
	default:
		return float64(this.minHeap.data[0]+this.maxHeap.data[0]) / 2
	}
}
