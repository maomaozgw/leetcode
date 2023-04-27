package p1552

import "sort"

func check(dist int, position []int, m int) bool {
	var pre = position[0]
	var cnt = 1
	for i := 1; i < len(position); i++ {
		if position[i]-pre >= dist {
			pre = position[i]
			cnt++
		}
	}
	return cnt >= m
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	var result = 0
	var left, right = 1, position[len(position)-1] - position[0]
	for left <= right {
		var mid = (left + right) >> 1
		if check(mid, position, m) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
