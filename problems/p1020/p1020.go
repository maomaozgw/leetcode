package p1020

var (
	nextX = []int{0, 1, 0, -1}
	nextY = []int{1, 0, -1, 0}
)

const (
	xOffset = 10000
)

func numEnclaves(grid [][]int) int {
	var result = 0
	var visited = make([][]bool, len(grid))
	var rl, cl = len(grid), len(grid[0])
	for i := range grid {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 || visited[i][j] {
				continue
			}
			var isClosed = true
			var surrounders []int
			surrounders = append(surrounders, i*xOffset+j)
			var islandCount = 0
			for len(surrounders) > 0 {
				var newSurrounders []int
				for _, v := range surrounders {
					y := v % xOffset
					x := v / xOffset
					visited[x][y] = true
					if x == 0 || y == 0 || x == rl-1 || y == cl-1 {
						isClosed = false
					}
					islandCount++
					for k := 0; k < 4; k++ {
						xN := x + nextX[k]
						yN := y + nextY[k]
						if xN < 0 || yN < 0 || xN == rl || yN == cl || grid[xN][yN] == 0 || visited[xN][yN] {
							continue
						}
						visited[xN][yN] = true
						newSurrounders = append(newSurrounders, xN*xOffset+yN)
					}
					surrounders = newSurrounders
				}
			}

			if isClosed {
				result += islandCount
			}
		}
	}
	return result
}
