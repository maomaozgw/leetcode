package p1926

const (
	wall   = '+'
	way    = '.'
	offset = 1000
)

var (
	nextX = []int{-1, 0, 0, 1}
	nextY = []int{0, -1, 1, 0}
)

func nearestExit(maze [][]byte, entrance []int) int {
	var m, n = len(maze), len(maze[0])
	var cache = make([][]int, m)
	for i := 0; i < len(maze); i++ {
		cache[i] = make([]int, n)
	}
	var q []int
	var currentX, currentY int
	q = append(q, entrance[0]*offset+entrance[1])
	steps := 0
	for len(q) > 0 {
		var newQ []int
		for i := 0; i < len(q); i++ {
			currentX, currentY = q[i]/offset, q[i]%offset
			if cache[currentX][currentY] > 0 {
				continue
			}
			cache[currentX][currentY] = steps
			if currentX == 0 || currentX == m-1 || currentY == 0 || currentY == n-1 {
				if currentX != entrance[0] || currentY != entrance[1] {
					return int(steps)
				}
			}
			for j := 0; j < 4; j++ {
				nx, ny := currentX+nextX[j], currentY+nextY[j]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && maze[nx][ny] == way {
					newQ = append(newQ, nx*offset+ny)
				}
			}
		}
		steps++
		q = newQ
	}
	return -1
}
