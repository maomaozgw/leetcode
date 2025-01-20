package p2661

const (
	offset = 100000
)

func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	rows := make([]int, m)
	cols := make([]int, n)
	pos := make([]int, m*n+1)
	for i, row := range mat {
		for j, v := range row {
			pos[v] = i*offset + j
		}
	}
	for i, v := range arr {
		p := pos[v]
		x := p / offset
		y := p % offset
		rows[x]++
		cols[y]++
		if rows[x] == n || cols[y] == m {
			return i
		}
	}
	return -1
}
