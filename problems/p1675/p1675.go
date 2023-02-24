package p1675

import (
	"container/heap"
	"math"
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

func minimumDeviation(nums []int) int {
	res := math.MaxInt
	minVal := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			nums[i] *= 2
		}
		minVal = min(minVal, nums[i])
	}
	h := Heap(nums)
	heap.Init(&h)
	for nums[0]%2 == 0 {
		v := heap.Pop(&h).(int) / 2
		minVal = min(v, minVal)
		heap.Push(&h, v)
		res = min(res, (nums[0] - minVal))
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
