package p1630

import (
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var n = len(l)
	var result = make([]bool, n)
	for i := 0; i < n; i++ {
		var subL = r[i] - l[i] + 1
		var subarr = make([]int, subL)
		copy(subarr, nums[l[i]:r[i]+1])
		result[i] = isArithmetic(subarr)
	}
	return result
}

func isArithmetic(arr []int) bool {
	sort.Ints(arr)
	var offset = arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != offset {
			return false
		}
	}
	return true

}
