// https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/

package p2279

import (
	"container/heap"
	"sort"
)

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

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i := range capacity {
		capacity[i] = capacity[i] - rocks[i]
	}
	sort.Ints(capacity)
	var cnt = 0
	for i := range capacity {
		if additionalRocks < capacity[i] {
			return cnt
		}
		additionalRocks -= capacity[i]
		cnt++
	}
	return cnt
}

func maximumBagsHeap(capacity []int, rocks []int, additionalRocks int) int {
	h := &Heap{
		lessFunc: func(i, j int) bool {
			return i < j
		},
	}
	heap.Init(h)
	var cnt = 0
	for idx := range capacity {
		remain := capacity[idx] - rocks[idx]
		if remain == 0 {
			cnt++
			continue
		}
		heap.Push(h, remain)
	}

	for additionalRocks > 0 && h.Len() > 0 {
		v := heap.Pop(h).(int)
		if additionalRocks < v {
			break
		}
		additionalRocks -= v
		cnt++
	}
	return cnt
}
