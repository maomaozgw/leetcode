package p153

func findMin(nums []int) int {
	idx := findK(nums, 0, len(nums)-1)
	return nums[(idx+1)%len(nums)]
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
