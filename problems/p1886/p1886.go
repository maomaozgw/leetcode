package p1886

func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		if equals(mat, target) {
			return true
		}
		rotate(mat)
	}
	return false
}

func rotate(matrix [][]int) {
	n := len(matrix)
	l := n / 2
	for i := 0; i < l; i++ {
		var (
			x1, y1 = 0 + i, 0 + i
			x2, y2 = 0 + i, n - 1 - i
			x3, y3 = n - 1 - i, n - 1 - i
			x4, y4 = n - 1 - i, 0 + i
		)
		for y1 < n-i-1 {
			matrix[x1][y1], matrix[x4][y4], matrix[x3][y3], matrix[x2][y2] = matrix[x4][y4], matrix[x3][y3], matrix[x2][y2], matrix[x1][y1]
			y1++
			x2++
			y3--
			x4--
		}
	}
}

func equals(mat, target [][]int) bool {
	var n = len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}

	return true
}
