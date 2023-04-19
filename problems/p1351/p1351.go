package p1351

func countNegatives(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	if grid[m-1][n-1] >= 0 {
		return 0
	}
	if grid[0][0] < 0 {
		return m * n
	}
	var result = 0
	var left, right = 0, m - 1

	for left < right {
		var mid = (left + right) >> 1
		if grid[mid][n-1] >= 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	for i := left; i < m; i++ {
		if grid[i][n-1] >= 0 {
			continue
		}
		if grid[i][0] < 0 {
			result += (m - i) * n
			break
		}
		var left, right = 0, n - 1
		for left < right {
			var mid = (left + right) >> 1
			if grid[i][mid] >= 0 {
				left = mid + 1
			} else {
				right = mid
			}
		}
		result += n - left
	}
	return result
}
