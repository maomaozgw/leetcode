package p452

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	currentEnd := -1<<31 - 1
	shotCnt := 0
	for i := range points {
		if points[i][0] > currentEnd {
			currentEnd = points[i][1]
			shotCnt++
		}
		currentEnd = min(currentEnd, points[i][1])
	}
	return shotCnt
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
