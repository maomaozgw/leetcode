package p912

func sortArray(nums []int) []int {
	fastSort(nums, 0, len(nums)-1)
	return nums
}

func fastSort(nums []int, leftBond, rightBond int) {
	if leftBond >= rightBond {
		return
	}
	var middle = nums[leftBond]
	var left, right = leftBond, rightBond
	for left != right {
		for left < right && nums[right] >= middle {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
		for left < right && nums[left] <= middle {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	fastSort(nums, left+1, rightBond)
	fastSort(nums, leftBond, left-1)
}
