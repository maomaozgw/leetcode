package p2530

import "container/heap"

type h struct {
	arr []int
}

func (h *h) Len() int {
	return len(h.arr)
}

func (h *h) Less(i, j int) bool {
	return h.arr[i] > h.arr[j]
}

func (h *h) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *h) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}

func (h *h) Pop() interface{} {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	res := int64(0)
	hp := &h{arr: make([]int, 0)}
	hp.arr = nums
	heap.Init(hp)
	for i := 0; i < k; i++ {
		top := heap.Pop(hp).(int)
		res += int64(top)
		top = (top + 2) / 3
		heap.Push(hp, top)
	}
	return res
}
