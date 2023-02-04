package p1779

func nearestValidPoint(x int, y int, points [][]int) int {
	var minLength = 100000
	var minInx = -1
	for i := range points {
		if points[i][0] != x && points[i][1] != y {
			continue
		}
		current := abs(x-points[i][0]) + abs(y-points[i][1])
		if current < minLength {
			minLength = current
			minInx = i
		}
	}
	return minInx
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
