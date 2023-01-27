package p54

func spiralOrder(matrix [][]int) []int {
	var left, right, top, bottom = 0, len(matrix[0]) - 1, 0, len(matrix) - 1
	var result []int
	var cnt = len(matrix) * len(matrix[0])

	for left <= right && top <= bottom && cnt > 0 {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
			cnt--
		}
		top++
		if cnt == 0 {
			break
		}
		for j := top; j <= bottom; j++ {
			result = append(result, matrix[j][right])
			cnt--
		}
		if cnt == 0 {
			break
		}
		right--
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
			cnt--
		}
		if cnt == 0 {
			break
		}
		bottom--
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
			cnt--
		}
		if cnt == 0 {
			break
		}
		left++
	}
	return result
}
