package p704

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {
	middle := left + (right-left)/2
	if middle == left {
		if nums[middle] == target {
			return middle
		}
		if nums[right] == target {
			return right
		}
		return -1
	}
	if nums[middle] > target {
		return binarySearch(nums, target, left, middle)
	} else if nums[middle] == target {
		return middle
	} else {
		return binarySearch(nums, target, middle, right)
	}
}
