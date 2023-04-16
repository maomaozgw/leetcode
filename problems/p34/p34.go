// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package p34

const (
	notFound = -1
)

func searchRange(nums []int, target int) []int {
	var (
		firstIndex int
		lastIndex  int
	)
	if len(nums) == 0 {
		return []int{notFound, notFound}
	}

	if nums[0] > target {
		return []int{notFound, notFound}
	}
	if nums[len(nums)-1] < target {
		return []int{notFound, notFound}
	}

	lastIndex = binarySearch(nums, 0, len(nums)-1, func(idx int) bool {
		return nums[idx] >= target && (idx == len(nums)-1 || nums[idx+1] > target)
	})
	if nums[lastIndex] != target {
		return []int{notFound, notFound}
	}
	firstIndex = binarySearch(nums, 0, lastIndex, func(idx int) bool {
		return nums[idx] >= target && (idx == 0 || nums[idx-1] <= target)
	})

	return []int{firstIndex, lastIndex}
}

func binarySearch(nums []int, left, right int, f func(idx int) bool) int {
	for left < right {
		var mid = (left + right) / 2
		if f(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
