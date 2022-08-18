// https://leetcode.com/problems/reduce-array-size-to-the-half/

package p1338

import "sort"

func minSetSize(arr []int) int {
	l := len(arr)
	h := l / 2
	eleCount := map[int]int{}
	for i := 0; i < l; i++ {
		eleCount[arr[i]]++
	}
	var eleCountSlice []int
	for _, val := range eleCount {
		eleCountSlice = append(eleCountSlice, val)
	}
	sort.Ints(eleCountSlice)
	var result = 0
	for i := len(eleCountSlice) - 1; i >= 0; i-- {
		l -= eleCountSlice[i]
		result += 1
		if l <= h {
			return result
		}
	}
	return -1
}
