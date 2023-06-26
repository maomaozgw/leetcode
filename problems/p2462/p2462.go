package p2462

import (
	"math"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func totalCost(costs []int, k int, candidates int) int64 {
	var result int64
	q1 := pq.NewWith(utils.IntComparator)
	q2 := pq.NewWith(utils.IntComparator)
	var cnt, i, j int
	j = len(costs) - 1
	for cnt < k {
		for q1.Size() < candidates && i <= j {
			q1.Enqueue(costs[i])
			i++
		}
		for q2.Size() < candidates && j >= i {
			q2.Enqueue(costs[j])
			j--
		}
		var c1, c2 = math.MaxInt, math.MaxInt
		if v1, ok := q1.Peek(); ok {
			c1 = v1.(int)
		}
		if v, ok := q2.Peek(); ok {
			c2 = v.(int)
		}
		if c1 <= c2 {
			result += int64(c1)
			q1.Dequeue()
		} else {
			result += int64(c2)
			q2.Dequeue()
		}
		cnt++
	}
	return result
}
