package p1578

import "container/heap"

type costs []int

func (c *costs) Len() int {
	return len(*c)
}

func (c *costs) Less(i, j int) bool {
	return (*c)[i] < (*c)[j]
}

func (c *costs) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *costs) Push(x any) {
	*c = append(*c, x.(int))
}

func (c *costs) Pop() any {
	old := *c
	n := len(old)
	item := old[n-1]
	*c = old[0 : n-1]
	return item
}

func minCost(colors string, neededTime []int) int {
	var total int
	var q costs
	var prev byte = colors[0]
	heap.Init(&q)
	heap.Push(&q, neededTime[0])
	for i := 1; i < len(colors); i++ {
		if colors[i] == prev {
			heap.Push(&q, neededTime[i])
		} else {
			prev = colors[i]
			for q.Len() > 1 {
				val := heap.Pop(&q)
				total += val.(int)
			}
			heap.Pop(&q)
			heap.Push(&q, neededTime[i])
		}
	}
	for q.Len() > 1 {
		val := heap.Pop(&q)
		total += val.(int)
	}
	return total
}
