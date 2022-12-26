package p55

func canJump(nums []int) bool {
	var l = len(nums)
	var dp = make([]bool, l)
	for i := l - 1; i > -1; i-- {
		maxL := i + nums[i]
		if maxL >= l-1 {
			dp[i] = true
			continue
		}
		for j := i + 1; j <= maxL; j++ {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}
