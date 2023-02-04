package p33

func search(nums []int, target int) int {
	if nums[0] < nums[len(nums)-1] {
		return binarySearch(nums, 0, len(nums)-1, target)
	}
	k := findK(nums, 0, len(nums)-1)
	if target >= nums[0] {
		return binarySearch(nums, 0, k, target)
	}
	return binarySearch(nums, k, len(nums)-1, target)

}

func findK(nums []int, left, right int) int {
	middle := (right + left) / 2
	if middle == left {
		if nums[middle] > nums[right] {
			return middle
		}
		return right
	}
	if nums[middle] > nums[middle+1] {
		return middle
	}
	if nums[middle] < nums[0] {
		return findK(nums, left, middle)
	}
	return findK(nums, middle, right)

}

func binarySearch(nums []int, left, right, target int) int {
	middle := (right + left) / 2
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
		return binarySearch(nums, left, middle, target)
	} else if nums[middle] == target {
		return middle
	}
	return binarySearch(nums, middle, right, target)

}
