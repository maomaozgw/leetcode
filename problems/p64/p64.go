package p64

func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])

	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			}
			if i != 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			}
			if i != 0 && j != 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
