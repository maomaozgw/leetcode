package p881

import "sort"

func numRescueBoats(people []int, limit int) int {
	var left, right = 0, len(people) - 1
	var result = 0
	sort.Ints(people)
	for left <= right {
		if people[left]+people[right] > limit {
			right--
		} else {
			left++
			right--
		}
		result++
	}
	return result
}
