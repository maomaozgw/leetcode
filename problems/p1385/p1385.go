package p1385

import (
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	var result = 0
	for i := range arr1 {
		var idx = binarySearch(arr2, arr1[i]-d)
		if idx < len(arr2) && arr2[idx] >= arr1[i]-d && arr2[idx] <= arr1[i]+d {
			continue
		}
		result++
	}
	return result
}

func binarySearch(nums []int, target int) int {
	var left, right = 0, len(nums)
	for left < right {
		var mid = (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
