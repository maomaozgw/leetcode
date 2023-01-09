package p973

import "container/heap"

type Heap [][]int

func (h *Heap) Less(i int, j int) bool {
	a, b := (*h)[i], (*h)[j]
	av := a[0]*a[0] + a[1]*a[1]
	bv := b[0]*b[0] + b[1]*b[1]
	return av < bv
}

func (h *Heap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	if k == len(points) {
		return points
	}
	if k == 0 {
		return [][]int{}
	}
	var h = &Heap{}
	heap.Init(h)
	for i := 0; i < len(points); i++ {
		heap.Push(h, points[i])
	}
	var result [][]int
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).([]int))
	}
	return result
}
