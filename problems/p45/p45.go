package p45

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var dp = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= 0 {
			continue
		}
		if i+nums[i] >= len(nums)-1 {
			dp[i] = 1
			continue
		}
		for j := 1; j <= nums[i]; j++ {
			if dp[i+j] == 0 {
				continue
			}
			if dp[i] == 0 {
				dp[i] = dp[i+j] + 1
			} else {
				dp[i] = min(dp[j+i]+1, dp[i])
			}
		}
	}
	return dp[0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
