package p2482

func onesMinusZeros(grid [][]int) [][]int {
	var rl, cl = len(grid), len(grid[0])
	var cols = make([]int, cl)
	var rows = make([]int, rl)
	for i := range grid {
		for j := range grid[i] {
			rows[i] += grid[i][j]
			cols[j] += grid[i][j]
		}
	}
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 2*(rows[i]+cols[j]) - rl - cl
		}
	}
	return grid
}
