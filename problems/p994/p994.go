package p994

const (
	offset = 100
)

var (
	nextX = []int{-1, 0, 1, 0}
	nextY = []int{0, 1, 0, -1}
)

func orangesRotting(grid [][]int) int {
	var q []int
	var goodCnt = 0
	var m, n = len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, i*offset+j)
			} else if grid[i][j] == 1 {
				goodCnt++
			}
		}
	}
	var timeCnt = 0
	for len(q) > 0 {
		var newQ []int
		for i := 0; i < len(q); i++ {
			cX, cY := q[i]/offset, q[i]%offset
			for j := 0; j < 4; j++ {
				nextX, nextY := cX+nextX[j], cY+nextY[j]
				if nextX < 0 || nextX == m || nextY < 0 || nextY == n || grid[nextX][nextY] != 1 {
					continue
				}
				grid[nextX][nextY] = 2
				goodCnt--
				newQ = append(newQ, nextX*offset+nextY)
			}
		}
		if len(newQ) > 0 {
			timeCnt++
		}
		q = newQ
	}
	if goodCnt == 0 {
		return timeCnt
	}
	return -1
}
