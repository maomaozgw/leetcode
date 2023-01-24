package p120

func minimumTotal(triangle [][]int) int {
	var l = len(triangle)
	var dp = make([]int, len(triangle[l-1]))
	for i := 0; i < len(triangle[l-1]); i++ {
		dp[i] = triangle[l-1][i]
	}
	for i := len(triangle) - 2; i > -1; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(triangle[i][j]+dp[j], triangle[i][j]+dp[j+1])
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
