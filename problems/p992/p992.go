package p992

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithLessKDistinct(nums, k) - subarraysWithLessKDistinct(nums, k-1)
}

func subarraysWithLessKDistinct(nums []int, k int) int {
	var left, right = 0, 0
	var numberCntMap = map[int]int{}
	var result = 0
	for ; right < len(nums); right++ {
		numberCntMap[nums[right]]++
		for left <= right && len(numberCntMap) > k {
			numberCntMap[nums[left]]--
			if numberCntMap[nums[left]] == 0 {
				delete(numberCntMap, nums[left])
			}
			left++
		}
		result += right - left + 1
	}
	return result
}
