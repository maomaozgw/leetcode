package p1502

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	var diff = arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
