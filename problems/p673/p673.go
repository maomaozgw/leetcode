package p673

func findNumberOfLIS(nums []int) int {
	var largest int
	var cntMap = make([]int, len(nums))
	var dp = make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
		cntMap[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				cntMap[i] = cntMap[j]
			} else if nums[i] > nums[j] && dp[i] == dp[j]+1 {
				cntMap[i] += cntMap[j]
			}
		}
		largest = max(largest, dp[i])
	}
	var result int
	for i := range dp {
		if dp[i] == largest {
			result += cntMap[i]
		}
	}

	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
