package p576

const (
	modVal = 1000000007
)

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var dp [][][]int
	dp = make([][][]int, m)
	for x := 0; x < m; x++ {
		dp[x] = make([][]int, n)
		for y := 0; y < n; y++ {
			dp[x][y] = make([]int, maxMove+1)
			for z := 0; z <= maxMove; z++ {
				dp[x][y][z] = -1
			}
		}
	}
	var dfs func(x, y, moves int) int
	dfs = func(x, y, moves int) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			// outside
			return 1
		}
		if moves <= 0 {
			return 0
		}
		if dp[x][y][moves] != -1 {
			return dp[x][y][moves]
		}
		var ans int
		ans += dfs(x-1, y, moves-1) % modVal
		ans += dfs(x+1, y, moves-1) % modVal
		ans += dfs(x, y-1, moves-1) % modVal
		ans += dfs(x, y+1, moves-1) % modVal
		dp[x][y][moves] = ans % modVal
		return dp[x][y][moves]
	}

	return dfs(startRow, startColumn, maxMove)
}
