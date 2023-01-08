package p347

import "container/heap"

type pair struct {
	val int
	cnt int
}

type Heap []pair

func (h *Heap) Less(i int, j int) bool {
	return (*h)[i].cnt > (*h)[j].cnt
}

func (h *Heap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(pair))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	var cntMap = make(map[int]int, 20001)
	for i := range nums {
		cntMap[nums[i]]++
	}
	h := &Heap{}
	heap.Init(h)
	for val, cnt := range cntMap {
		heap.Push(h, pair{
			val: val,
			cnt: cnt,
		})
	}
	var result []int
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(pair)
		result = append(result, p.val)
	}
	return result
}
