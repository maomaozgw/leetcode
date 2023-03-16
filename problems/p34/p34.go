// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package p34

const (
	notFound = -1
)

func searchRange(nums []int, target int) []int {

	var (
		firstIndex = notFound
		lastIndex  = notFound
		valueIndex = notFound
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

	if nums[0] == target {
		valueIndex = 0
	} else {
		valueIndex = binarySearch(nums, target, 0, len(nums)-1)
	}

	if valueIndex == notFound {
		return []int{notFound, notFound}
	}
	for firstIndex = valueIndex; firstIndex >= 0; firstIndex-- {
		if nums[firstIndex] != target {
			break
		}
	}
	firstIndex++
	for lastIndex = valueIndex; lastIndex < len(nums); lastIndex++ {
		if nums[lastIndex] != target {
			break
		}
	}
	lastIndex--
	return []int{firstIndex, lastIndex}
}

func binarySearch(nums []int, target int, left, right int) int {
	middle := left + (right-left)/2
	if middle == left {
		if nums[middle] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		return notFound
	}
	if nums[middle] > target {
		return binarySearch(nums, target, left, middle)
	} else if nums[middle] == target {
		return middle
	} else {
		return binarySearch(nums, target, middle, right)
	}
}
