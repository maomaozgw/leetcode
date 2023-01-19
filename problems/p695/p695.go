package p695

func maxAreaOfIsland(grid [][]int) int {
	const (
		xOffset = 1000
	)
	var (
		nextX = []int{-1, 0, 1, 0}
		nextY = []int{0, 1, 0, -1}
		rl    = len(grid)
		cl    = len(grid[0])
	)

	result := grid[0][0]
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			if grid[i][j] == 0 {
				continue
			}
			var currentSize = 1
			var surrounders []int
			grid[i][j] = 0
			surrounders = append(surrounders, i*xOffset+j)
			for len(surrounders) > 0 {
				var newSurrounders []int
				for _, v := range surrounders {
					y := v % xOffset
					x := v / xOffset
					for k := 0; k < 4; k++ {
						xN := x + nextX[k]
						yN := y + nextY[k]
						if xN < 0 || yN < 0 || xN == rl || yN == cl || grid[xN][yN] == 0 {
							continue
						}
						grid[xN][yN] = 0
						currentSize++
						newSurrounders = append(newSurrounders, xN*xOffset+yN)
					}
					surrounders = newSurrounders

				}
			}
			result = max(result, currentSize)
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
