package p540

func singleNonDuplicate(nums []int) int {
	return binarySearch(nums, 0, len(nums)-1)
}

func binarySearch(nums []int, left, right int) int {
	var middle = (left + right) / 2
	if middle == left {
		if middle == 0 {
			return nums[middle]
		}
		if nums[middle] == nums[middle-1] {
			return nums[middle+1]
		}
		return nums[middle]
	}
	if nums[middle] == nums[middle-1] {
		if middle%2 == 1 {
			return binarySearch(nums, middle, right)
		}
		return binarySearch(nums, left, middle)
	} else if nums[middle] == nums[middle+1] {
		if middle%2 == 1 {
			return binarySearch(nums, left, middle)
		}
		return binarySearch(nums, middle, right)
	}
	return nums[middle]
}
