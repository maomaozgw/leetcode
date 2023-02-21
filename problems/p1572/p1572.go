package p1572

func diagonalSum(mat [][]int) int {
	var n = len(mat)
	var result = 0
	for i := 0; i < n; i++ {
		result += mat[i][i]
		result += mat[i][n-1-i]
	}
	if n%2 == 0 {
		return result
	}
	result -= mat[n/2][n/2]
	return result
}
