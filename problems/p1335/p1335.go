package p1335

var (
	dp  = make([][]int, 301)
	row = make([]int, 11)
)

func init() {
	for i := 0; i < 11; i++ {
		row[i] = -1
	}
}

func minDifficulty(jobDifficulty []int, d int) int {
	var n = len(jobDifficulty)
	if n < d {
		return -1
	}

	for i := 0; i < 301; i++ {
		dp[i] = make([]int, 11)
		copy(dp[i], row)
	}
	return solve(jobDifficulty, n, 0, d)
}

func solve(jobs []int, n, idx, d int) int {
	if d == 1 {
		var max = 0
		for i := idx; i < n; i++ {
			if jobs[i] > max {
				max = jobs[i]
			}
		}
		return max
	}
	if dp[idx][d] != -1 {
		return dp[idx][d]
	}
	var maxVal, result = -1, 10000
	for i := idx; i <= n-d; i++ {
		maxVal = max(maxVal, jobs[i])
		result = min(result, maxVal+solve(jobs, n, i+1, d-1))
	}
	dp[idx][d] = result
	return result
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
