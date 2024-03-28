package p2958

func maxSubarrayLength(nums []int, k int) int {
	var result = 0
	var cntMap = map[int]int{}
	var left = 0
	for right := 0; right < len(nums); right++ {
		cntMap[nums[right]]++
		for cntMap[nums[right]] > k {
			cntMap[nums[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
