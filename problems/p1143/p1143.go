// https://leetcode.com/problems/longest-common-subsequence/description/

package p1143

func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	var dp = make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			c1 := text1[i]
			c2 := text2[j]
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
