package p1356

import (
	"sort"
)

func sortByBits(arr []int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		ib, jb := get1Bits(arr[i]), get1Bits(arr[j])
		if ib == jb {
			return arr[i] < arr[j]
		}
		return ib < jb
	})
	return arr
}

func get1Bits(num int) int {
	c := 0
	for num != 0 {
		num = num & (num - 1)
		c++
	}
	return c
}
