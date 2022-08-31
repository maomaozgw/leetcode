// https://leetcode.com/problems/pacific-atlantic-water-flow/

package p417

const (
	offset = 1000
)

const (
	canP = 1 << iota
	canA

	canAll = canP | canA
)

var (
	nextX = []int{-1, 0, 1, 0}
	nextY = []int{0, 1, 0, -1}
)

func pacificAtlantic(heights [][]int) [][]int {
	var marks [][]int
	m := len(heights)
	n := len(heights[0])
	marks = make([][]int, m)
	var result [][]int
	for i := 0; i < m; i++ {
		marks[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		marks[i][0] |= canP
		if i < m-1 {
			dfs(i, 0, m, n, marks, heights)
		}
		marks[i][n-1] |= canA
		if i > 0 {
			dfs(i, n-1, m, n, marks, heights)
		}

	}
	for i := 1; i < n; i++ {
		marks[0][i] |= canP
		dfs(0, i, m, n, marks, heights)
		marks[m-1][i-1] |= canA
		dfs(m-1, i-1, m, n, marks, heights)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if marks[i][j] != canAll {
				continue
			}
			result = append(result, []int{i, j})
		}
	}
	return result
}

func dfs(i, j, m, n int, marks [][]int, heights [][]int) {
	var checkSlice []int
	checkSlice = append(checkSlice, i*offset+j)
	for len(checkSlice) > 0 {
		var newCheckSlice []int
		for _, v := range checkSlice {
			x, y := v/offset, v%offset
			v = heights[x][y]
			for k := 0; k < 4; k++ {
				xN := nextX[k] + x
				yN := nextY[k] + y
				if xN < 0 || yN < 0 || xN == m || yN == n || marks[xN][yN]&marks[x][y] == marks[x][y] || heights[xN][yN] < v {
					continue
				}
				marks[xN][yN] |= marks[x][y]
				newCheckSlice = append(newCheckSlice, xN*offset+yN)
			}
		}
		checkSlice = newCheckSlice
	}
}
