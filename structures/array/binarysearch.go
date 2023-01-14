package array

import "golang.org/x/exp/constraints"

const (
	notFound = -1
)

func BinarySearch[T constraints.Integer | constraints.Float](nums []T, target T, left, right int) int {
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
		return BinarySearch(nums, target, left, middle)
	} else if nums[middle] == target {
		return middle
	} else {
		return BinarySearch(nums, target, middle, right)
	}
}

func BinarySearchL(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		pos := left + (right-left)/2
		val := nums[pos]
		if val == target {
			return pos
		} else if val < target {
			left = pos + 1
		} else /* if val > target */ {
			right = pos - 1
		}
	}
	return -1
}
