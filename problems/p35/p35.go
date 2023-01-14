package p35

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums))
}

func binarySearch(nums []int, target, left, right int) int {
	middle := (right + left) / 2
	if middle == left {
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			return middle + 1
		} else {
			return middle
		}
	}
	if nums[middle] > target {
		return binarySearch(nums, target, left, middle)
	} else if nums[middle] < target {
		return binarySearch(nums, target, middle, right)
	} else {
		return middle
	}
}
