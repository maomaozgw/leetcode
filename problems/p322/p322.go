package p322

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c < 0 {
				continue
			}
			dp[a] = min(dp[a-c]+1, dp[a])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
