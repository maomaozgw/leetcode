package p1861

func rotateTheBox(box [][]byte) [][]byte {

	var m = len(box)
	var n = len(box[0])
	var resultRow = make([]byte, m)
	for i := range resultRow {
		resultRow[i] = '.'
	}
	var result = make([][]byte, n)
	for i := range result {
		result[i] = make([]byte, m)
		copy(result[i], resultRow)
	}
	for i := 0; i < m; i++ {
		var row = box[i]
		var lastWriteAbleIdx = n - 1
		for j := n - 1; j >= 0; j-- {
			switch row[j] {
			case '#':
				row[j] = '.'
				result[lastWriteAbleIdx][m-1-i] = '#'
				lastWriteAbleIdx--
			case '*':
				result[j][m-1-i] = '*'
				lastWriteAbleIdx = j - 1
			}
		}
	}
	return result
}
