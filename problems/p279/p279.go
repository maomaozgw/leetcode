// https://leetcode.com/problems/perfect-squares/

package p279

const (
	maxInt32 = 1<<31 - 1
	maxCount = 101
)

var (
	prefectSquares []int
)

func init() {
	prefectSquares = make([]int, maxCount)
	for i := 0; i < maxCount; i++ {
		prefectSquares[i] = i * i
	}
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = maxInt32
	}

	for _, val := range prefectSquares {
		if val == n {
			return 1
		}
		if val > n {
			break
		}
		dp[val] = 1
	}
	for i := 2; i <= n; i++ {
		if dp[i] == 1 {
			continue
		}

		for j := 1; j < maxCount; j++ {
			if i < prefectSquares[j] {
				break
			}
			dp[i] = min(dp[i], dp[i-prefectSquares[j]]+dp[prefectSquares[j]])
		}
	}
	return dp[n]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
