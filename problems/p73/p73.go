package p73

func setZeroes(matrix [][]int) {
	var m = len(matrix)
	var n = len(matrix[0])
	var zeroRows = make([]bool, m)
	var zeroCols = make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = true
				zeroCols[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if zeroRows[i] || zeroCols[j] {
				matrix[i][j] = 0
			}
		}
	}
}
