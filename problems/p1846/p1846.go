package p1846

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	var result = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > result {
			result++
		}
	}
	return result
}
