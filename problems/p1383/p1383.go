// https://leetcode.com/problems/maximum-performance-of-a-team/

package p1383

import (
	pq "github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
	"sort"
)

const modVal = 1000000007

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	var employees = make([][]int, n)
	for i := 0; i < n; i++ {
		employees[i] = []int{efficiency[i], speed[i]}
	}
	sort.Slice(employees, func(i, j int) bool {
		return employees[i][0] > employees[j][0]
	})
	var sum = 0
	var e = 0
	var performance = 0
	var q = pq.NewWith(utils.IntComparator)
	for i := 0; i < k; i++ {
		sum += employees[i][1]
		e = employees[i][0]
		q.Push(employees[i][1])
		performance = max(performance, sum*e)
	}
	for i := k; i < n; i++ {
		min, _ := q.Pop()
		sum -= min.(int)
		sum += employees[i][1]
		q.Push(employees[i][1])
		e = employees[i][0]
		performance = max(performance, sum*e)
	}
	return performance % modVal
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
