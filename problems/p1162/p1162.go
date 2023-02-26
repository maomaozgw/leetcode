package p1162

const (
	offset = 1000
)

var (
	nextX = []int{0, -1, 0, 1}
	nextY = []int{-1, 0, 1, 0}
)

func maxDistance(grid [][]int) int {
	var l = len(grid)
	var q []int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if grid[i][j] == 1 {
				q = append(q, i*offset+j)
			}
		}
	}
	var dist = 0
	if len(q) == 0 || len(q) == l*l {
		return -1
	}
	for len(q) > 0 {
		dist++
		var newQ []int
		for i := 0; i < len(q); i++ {
			var x, y = q[i] / offset, q[i] % offset
			for j := 0; j < 4; j++ {
				var nx, ny = x + nextX[j], y + nextY[j]
				if nx < 0 || ny < 0 || nx >= l || ny >= l || grid[nx][ny] != 0 {
					continue
				}
				grid[nx][ny] = 1
				newQ = append(newQ, nx*offset+ny)
			}
		}
		q = newQ
	}
	return dist - 1
}
