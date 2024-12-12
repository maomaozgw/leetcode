package p2558

import (
	"container/heap"
	"math"
)

type h struct {
	data []int
}

func (h *h) Len() int {
	return len(h.data)
}
func (h *h) Less(i, j int) bool {
	return h.data[i] > h.data[j]
}
func (h *h) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *h) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}
func (h *h) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	var h = h{data: gifts}
	heap.Init(&h)
	for i := 0; i < k; i++ {
		x := heap.Pop(&h).(int)
		if x == 1 {
			return int64(h.Len()) + 1
		}
		heap.Push(&h, int(math.Sqrt(float64(x))))
	}
	var result int64 = 0
	for i := range h.data {
		result += int64(h.data[i])
	}
	return result
}
