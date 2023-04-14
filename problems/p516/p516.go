package p516

func longestPalindromeSubseq(s string) int {
	var l = len(s)
	var dp = make([]int, l)
	for i := l - 1; i >= 0; i-- {
		var tmpDP = make([]int, l)
		tmpDP[i] = 1
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				tmpDP[j] = 2 + dp[j-1]
			} else {
				tmpDP[j] = max(dp[j], tmpDP[j-1])
			}
		}
		dp = tmpDP
	}
	return dp[l-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
