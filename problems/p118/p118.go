package p118

func generate(numRows int) [][]int {
	var result = make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
