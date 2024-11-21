package p2257

const (
	offset = 100000
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	var grid = make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, v := range walls {
		grid[v[0]][v[1]] = 2
	}
	var q []int
	for _, v := range guards {
		grid[v[0]][v[1]] = 3
		q = append(q, v[0]*offset+v[1])
	}

	for i := 0; i < len(q); i++ {
		v := q[i]
		x := v / offset
		y := v % offset
		for j := x + 1; j < m; j++ {
			if grid[j][y] >= 2 {
				break
			}
			grid[j][y] = 1
		}
		for j := x - 1; j >= 0; j-- {
			if grid[j][y] >= 2 {
				break
			}
			grid[j][y] = 1
		}
		for j := y + 1; j < n; j++ {
			if grid[x][j] >= 2 {
				break
			}
			grid[x][j] = 1
		}
		for j := y - 1; j >= 0; j-- {
			if grid[x][j] >= 2 {
				break
			}
			grid[x][j] = 1
		}
	}
	var result = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				result++
			}
		}
	}
	return result
}
