package p1091

const (
	offset = 1000
)

var (
	nextX = []int{0, 0, 1, 1, 1, -1, -1, -1}
	nextY = []int{1, -1, -1, 0, 1, -1, 0, 1}
)

func shortestPathBinaryMatrix(grid [][]int) int {
	var n = len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}
	var q []int
	var result = 1
	q = append(q, 0)
	for len(q) > 0 {
		var newQ []int
		for i := range q {
			cx, cy := q[i]/offset, q[i]%offset
			for j := 0; j < 8; j++ {
				nx, ny := cx+nextX[j], cy+nextY[j]
				if nx < 0 || ny < 0 || nx == n || ny == n || grid[nx][ny] == 1 {
					continue
				} else if n-nx == 1 && n-ny == 1 {
					return result + 1
				}
				newQ = append(newQ, nx*offset+ny)
				grid[nx][ny] = 1
			}
		}
		q = newQ
		result++
	}
	return -1
}
