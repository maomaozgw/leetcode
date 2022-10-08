// https://leetcode.com/problems/find-k-closest-elements/

package p658

import (
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	if k == len(arr) {
		return arr
	}
	var closet = make([]int, 20001)
	for i := 0; i < len(arr); i++ {
		closet[arr[i]+10000] = abs(arr[i] - x)
	}
	sort.Slice(arr, func(i, j int) bool {
		ci, cj := closet[arr[i]+10000], closet[arr[j]+10000]
		if ci == cj {
			return arr[i] < arr[j]
		}
		return ci < cj
	})

	var result = arr[:k]
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
