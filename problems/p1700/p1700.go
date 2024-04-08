package p1700

func countStudents(students []int, sandwiches []int) int {
	var cCnt, sCnt = 0, 0
	for _, num := range students {
		if num == 0 {
			cCnt++
		} else {
			sCnt++
		}
	}
	for _, n := range sandwiches {
		if n == 0 && cCnt > 0 {
			cCnt--
		} else if n == 1 && sCnt > 0 {
			sCnt--
		} else {
			break
		}
	}
	return cCnt + sCnt
}
