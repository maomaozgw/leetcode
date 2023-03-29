package p1402

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	var result int
	var lastSum = 0
	var total = 0
	for i := len(satisfaction) - 1; i > -1; i-- {
		total = total + lastSum + satisfaction[i]
		lastSum += satisfaction[i]
		if total >= result {
			result = total
		} else {
			return result
		}
	}
	return result
}
