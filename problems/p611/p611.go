package p611

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var result = 0
	for k := 2; k < len(nums); k++ {
		var left, right = 0, k - 1
		for left < right {
			if nums[left]+nums[right] > nums[k] {
				result += right - left
				right--
			} else {
				left++
			}
		}
	}
	return result
}
