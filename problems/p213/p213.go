package p213

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var result = nums[0]
	result = max(result, dp(nums[1:]))
	result = max(result, dp(nums[:len(nums)-1]))
	return result
}

func dp(nums []int) int {
	var first, second = 0, 0
	for idx := range nums {
		var t = max(nums[idx]+first, second)
		first = second
		second = t
	}
	return second
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
