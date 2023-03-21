package p130

const (
	offset = 1000
	mask   = '.'
	o      = 'O'
	x      = 'X'
)

var (
	nextX = []int{0, 1, 0, -1}
	nextY = []int{1, 0, -1, 0}
)

func solve(board [][]byte) {
	var q []int
	var m, n = len(board), len(board[0])
	for i := 0; i < n; i++ {
		if board[0][i] == o {
			q = append(q, i)
			board[0][i] = mask
		}
		if board[m-1][i] == o {
			q = append(q, (m-1)*offset+i)
			board[m-1][i] = mask
		}
	}
	for i := 1; i < m-1; i++ {
		if board[i][0] == o {
			q = append(q, i*offset)
			board[i][0] = mask
		}
		if board[i][n-1] == o {
			q = append(q, i*offset+(n-1))
			board[i][n-1] = mask
		}
	}
	for len(q) > 0 {
		var newQ []int
		for i := range q {
			cx, cy := q[i]/offset, q[i]%offset
			for i := 0; i < 4; i++ {
				nx, ny := cx+nextX[i], cy+nextY[i]
				if nx < 0 || ny < 0 || nx == m || ny == n || board[nx][ny] != o {
					continue
				}
				newQ = append(newQ, nx*offset+ny)
				board[nx][ny] = mask
			}
		}
		q = newQ
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == mask {
				board[i][j] = o
			} else {
				board[i][j] = x
			}
		}
	}
}
