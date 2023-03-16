package p986

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}
	var p1, p2 = 0, 0
	var result [][]int
	for p1 < len(firstList) && p2 < len(secondList) {
		if firstList[p1][1] < secondList[p2][0] {
			p1++
			continue
		} else if firstList[p1][0] > secondList[p2][1] {
			p2++
			continue
		}

		result = append(result, []int{max(firstList[p1][0], secondList[p2][0]), min(firstList[p1][1], secondList[p2][1])})

		if firstList[p1][1] < secondList[p2][1] {
			p1++
			continue
		} else {
			p2++
			continue
		}
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
