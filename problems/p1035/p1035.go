package p1035

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	var dp = make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			c1 := nums1[i]
			c2 := nums2[j]
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[l1][l2]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
