package p1253gan

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var result [][]int
	result = append(result, make([]int, len(colsum)))
	result = append(result, make([]int, len(colsum)))
	var oneList []int
	for idx, val := range colsum {
		if val == 2 {
			upper -= 1
			lower -= 1
			result[0][idx], result[1][idx] = 1, 1
			continue
		} else if val == 0 {
			continue
		}
		oneList = append(oneList, idx)
	}
	if upper < 0 || lower < 0 {
		// 必须分配的分配完了就不够了
		return [][]int{}
	}

	for _, idx := range oneList {
		if upper != 0 {
			upper -= 1
			result[0][idx] = 1
			continue
		}
		if lower != 0 {
			lower -= 1
			result[1][idx] = 1
			continue
		}
		return [][]int{}
	}
	if upper != 0 || lower != 0 {
		return [][]int{}
	}
	return result
}
