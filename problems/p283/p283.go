package p283

func moveZeroes(nums []int) {
	var left, right = 0, 0
	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			continue
		}
		nums[left] = nums[right]
		left++
	}
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}
