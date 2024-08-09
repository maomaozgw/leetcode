package p840

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 {
		return 0
	}
	var result = 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[0])-3; j++ {
			calc(grid, i, j, &result)
		}
	}
	return result
}

func calc(grid [][]int, i, j int, result *int) {
	var rowSum = make([]int, 3)
	var colSum = make([]int, 3)
	var diagSum = make([]int, 2)
	var set = make(map[int]bool)
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if grid[i+k][j+l] < 1 || grid[i+k][j+l] > 9 {
				return
			}
			set[grid[i+k][j+l]] = true
			rowSum[k] += grid[i+k][j+l]
			colSum[l] += grid[i+k][j+l]
			if k == l {
				diagSum[0] += grid[i+k][j+l]
			}
			if k+l == 2 {
				diagSum[1] += grid[i+k][j+l]
			}
		}
	}
	if len(set) != 9 {
		return
	}
	if rowSum[0] != rowSum[1] || rowSum[1] != rowSum[2] {
		return
	}
	if colSum[0] != colSum[1] || colSum[1] != colSum[2] {
		return
	}
	if diagSum[0] != diagSum[1] {
		return
	}
	*result++
}
