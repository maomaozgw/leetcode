package p304

type NumMatrix struct {
	data [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] = matrix[i][j] + matrix[i][j-1]
		}
	}
	return NumMatrix{
		data: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var result = 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			result += this.data[i][col2]
			continue
		}
		result += this.data[i][col2] - this.data[i][col1-1]
	}
	return result
}
