package p746

func minCostClimbingStairs(cost []int) int {
	var dp = cost
	for i := 2; i < len(cost); i++ {
		cost[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
