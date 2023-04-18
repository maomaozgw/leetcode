package p1608

import (
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	var left, right = 1, len(nums) + 1
	for left < right {
		var mid = (right + left) / 2
		var i = gatherCnt(nums, mid)
		if mid == i {
			return mid
		} else if i > mid {
			left = mid + 1
		} else if i < mid {
			right = mid
		}
	}
	return -1
}

func gatherCnt(nums []int, target int) int {
	var left, right = 0, len(nums) - 1

	for left <= right {
		var mid = (right + left) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return len(nums) - left
}
