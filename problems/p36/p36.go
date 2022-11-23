package p36

const (
	count = 9
	start = '1'
	empty = '.'
)

func newRecord() [][]int {
	var result = make([][]int, count)
	for i := range result {
		result[i] = make([]int, count)
	}
	return result
}

func isValidSudoku(board [][]byte) bool {
	var lineRecord, rowRecord, blockRecord = newRecord(), newRecord(), newRecord()
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			val := board[i][j]
			if val == empty {
				continue
			}
			bi := i/3*3 + j/3
			val -= start
			if lineRecord[i][val] > 0 || rowRecord[j][val] > 0 || blockRecord[bi][val] > 0 {
				return false
			}

			lineRecord[i][val]++
			rowRecord[j][val]++
			blockRecord[bi][val]++
		}
	}
	return true
}
