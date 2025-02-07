package p3160

func queryResults(limit int, queries [][]int) []int {
	var colorMap = map[int]int{}
	var result = make([]int, len(queries))
	var colorInfo = map[int]int{}
	for idx, q := range queries {
		oldColor := colorInfo[q[0]]
		colorMap[oldColor]--
		if colorMap[oldColor] == 0 {
			delete(colorMap, oldColor)
		}
		colorInfo[q[0]] = q[1]
		colorMap[q[1]]++
		result[idx] = len(colorMap) - 1
	}
	return result
}
