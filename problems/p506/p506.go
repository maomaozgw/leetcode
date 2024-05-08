package p506

import (
	"container/heap"
	"strconv"
)

type item struct {
	score int
	index int
}

type hl []item

func (h hl) Less(i, j int) bool {
	return h[i].score > h[j].score
}
func (h hl) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h hl) Len() int {
	return len(h)
}
func (h *hl) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *hl) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findRelativeRanks(score []int) []string {
	var result = make([]string, len(score))
	var data hl
	for i := 0; i < len(score); i++ {
		data = append(data, item{
			score: score[i],
			index: i,
		})
	}
	heap.Init(&data)
	hitem := heap.Pop(&data).(item)
	result[hitem.index] = "Gold Medal"
	if data.Len() == 0 {
		return result
	}
	hitem = heap.Pop(&data).(item)
	result[hitem.index] = "Silver Medal"
	if data.Len() == 0 {
		return result
	}
	hitem = heap.Pop(&data).(item)
	result[hitem.index] = "Bronze Medal"
	for i := 4; i <= len(score); i++ {
		hitem := heap.Pop(&data).(item)
		result[hitem.index] = strconv.Itoa(i)
	}

	return result
}
