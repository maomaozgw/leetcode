// https://leetcode.com/problems/number-of-islands/

package p200

func numIslands(grid [][]byte) int {
	const (
		xOffset = 1000
	)
	var (
		nextX = []int{-1, 0, 1, 0}
		nextY = []int{0, 1, 0, -1}
		rl    = len(grid)
		cl    = len(grid[0])
	)
	islandMark := make([][]bool, rl)
	for i := range islandMark {
		islandMark[i] = make([]bool, cl)
	}
	islandCount := 0
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			if grid[i][j] == '0' || islandMark[i][j] {
				continue
			}
			islandCount++
			var surrounders []int
			surrounders = append(surrounders, i*xOffset+j)
			for len(surrounders) > 0 {
				var newSurrounders []int
				for _, v := range surrounders {
					y := v % xOffset
					x := v / xOffset
					for k := 0; k < 4; k++ {
						xN := x + nextX[k]
						yN := y + nextY[k]
						if xN < 0 || yN < 0 || xN == rl || yN == cl || grid[xN][yN] == '0' || islandMark[xN][yN] {
							continue
						}
						islandMark[xN][yN] = true
						newSurrounders = append(newSurrounders, xN*xOffset+yN)
					}
					surrounders = newSurrounders

				}
			}
		}
	}
	return islandCount
}
