// https://leetcode.com/problems/non-overlapping-intervals/

package p435

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var end = -1 << 31
	var eraseCnt = 0
	for idx := 0; idx < len(intervals); idx++ {
		if intervals[idx][0] >= end {
			end = intervals[idx][1]
			continue
		}
		eraseCnt++
	}
	return eraseCnt
}
