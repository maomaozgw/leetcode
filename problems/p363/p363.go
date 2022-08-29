package p363

func maxSumSubmatrix(matrix [][]int, k int) int {
	rl := len(matrix)
	cl := len(matrix[0])
	sum := make([][]int, rl+1)
	for i := 0; i <= rl; i++ {
		sum[i] = make([]int, cl+1)
	}
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			sum[i+1][j+1] = matrix[i][j] + sum[i+1][j] + sum[i][j+1] - sum[i][j]
		}
	}
	result := -10000000000

	for i := 1; i <= rl; i++ {
		for j := 1; j <= cl; j++ {
			for i2 := i; i2 <= rl; i2++ {
				for j2 := j; j2 <= cl; j2++ {
					current := sum[i2][j2] - sum[i-1][j2] - sum[i2][j-1] + sum[i-1][j-1]
					if current == k {
						return k
					}
					if current < k && current > result {
						result = current
					}
				}
			}
		}
	}
	return result
}
