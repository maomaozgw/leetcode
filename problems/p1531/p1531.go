// https://leetcode.com/problems/string-compression-ii/

package p1531

func getLengthOfOptimalCompression(s string, k int) int {
	var n = len(s)
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = n
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if j > 0 {
				dp[i][j] = dp[i-1][j-1]
			}
			var removed, cnt int
			for p := i; p > 0; p-- {
				if s[p-1] == s[i-1] {
					cnt++
				} else {
					removed++
				}
				if removed > j {
					break
				}
				dp[i][j] = minInt(dp[i][j], dp[p-1][j-removed]+calcLen(cnt))
			}
		}
	}
	return dp[n][k]
}

func calcLen(len int) int {
	switch {
	case len == 0:
		return 0
	case len == 1:
		return 1
	case len < 10:
		return 2
	case len < 100:
		return 3
	default:
		return 4
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
