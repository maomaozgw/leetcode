package p1975

func maxMatrixSum(matrix [][]int) int64 {
	var negetiveCnt = 0
	var minNum = 1 << 31
	var sum = int64(0)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				negetiveCnt++
				sum += int64(-matrix[i][j])
				minNum = min(minNum, -matrix[i][j])
			} else {
				sum += int64(matrix[i][j])
				minNum = min(minNum, matrix[i][j])
			}
		}
	}
	if negetiveCnt%2 == 0 {
		return sum
	}
	return sum - int64(minNum*2)
}
