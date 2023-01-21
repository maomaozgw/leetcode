package p542

const (
	offset = 100000
)

var (
	nextX = []int{-1, 0, 1, 0}
	nextY = []int{0, 1, 0, -1}
)

func updateMatrix(mat [][]int) [][]int {
	var q []int
	var m, n = len(mat), len(mat[0])
	var result = make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				result[i][j] = 0
				q = append(q, i*offset+j)
			}
		}
	}
	var distance = 1
	for len(q) > 0 {
		var newQ []int
		for i := 0; i < len(q); i++ {
			currentX, currentY := q[i]/offset, q[i]%offset
			for j := 0; j < 4; j++ {
				nx, ny := currentX+nextX[j], currentY+nextY[j]
				if nx < 0 || nx == m || ny < 0 || ny == n || mat[nx][ny] == 0 {
					continue
				}
				mat[nx][ny] = 0
				result[nx][ny] = distance
				newQ = append(newQ, nx*offset+ny)
			}
		}
		q = newQ
		distance++
	}
	return result
}
