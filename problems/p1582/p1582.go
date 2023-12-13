package p1582

func numSpecial(mat [][]int) int {
	var result = 0
	for i := range mat {
		colIdx := checkRow(mat, i)
		if colIdx < 0 {
			continue
		}
		if checkColumn(mat, i, colIdx) {
			result++
		}
	}
	return result
}

func checkRow(mat [][]int, rowIdx int) int {
	var result = -1
	for colIdx := range mat[0] {
		if mat[rowIdx][colIdx] == 1 {
			if result >= 0 {
				return -1
			}
			result = colIdx
		}
	}
	return result
}

func checkColumn(mat [][]int, excludeRowIdx, colIdx int) bool {
	for rowIdx := range mat {
		if mat[rowIdx][colIdx] == 1 && rowIdx != excludeRowIdx {
			return false
		}
	}
	return true
}
