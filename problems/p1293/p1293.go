// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/

package p1293

var (
	nextX = []int{-1, 0, 0, 1}
	nextY = []int{0, -1, 1, 0}
)

type item struct {
	x, y, k int
}

func shortestPath(grid [][]int, k int) int {
	var m = len(grid)
	var n = len(grid[0])
	var remainMap = make([][]int, m)
	var remainLine = make([]int, n)
	for i := 0; i < n; i++ {
		remainLine[i] = -1
	}
	for i := 0; i < m; i++ {
		remainMap[i] = make([]int, n)
		copy(remainMap[i], remainLine)
	}
	var steps = 0
	var q []item
	q = append(q, item{0, 0, k})
	for len(q) > 0 {
		var newQ []item
		for i := 0; i < len(q); i++ {
			current := q[i]
			if current.x == m-1 && current.y == n-1 {
				return steps
			}
			for j := 0; j < 4; j++ {
				nx, ny := current.x+nextX[j], current.y+nextY[j]
				if nx < 0 || ny < 0 || nx >= m || ny >= n {
					continue
				}
				var remain = current.k - grid[nx][ny]
				if remain >= 0 && remainMap[nx][ny] < remain {
					newQ = append(newQ, item{nx, ny, remain})
					remainMap[nx][ny] = remain
				}
			}
		}
		q = newQ
		steps++
	}
	return -1
}
