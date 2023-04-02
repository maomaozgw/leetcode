package p2300

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	var p = len(potions)
	sort.Ints(potions)
	for i := range spells {
		var s = spells[i]
		var left, right = 0, p - 1
		for left <= right {
			var mid = (left + right) / 2
			var v int64 = int64(s) * int64(potions[mid])
			if v >= success {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		spells[i] = p - left
	}
	return spells
}
