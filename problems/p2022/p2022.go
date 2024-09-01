package p2022

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	var result = make([][]int, m)
	var k = 0
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		copy(result[i], original[k:k+n])
		k += n
	}
	return result
}

func construct2DArrayOld(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	var result = make([][]int, m)
	var k = 0
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = original[k]
			k++
		}
	}
	return result
}
