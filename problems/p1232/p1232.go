package p1232

func checkStraightLine(coordinates [][]int) bool {
	var lastRatio = getRatio(coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1])
	for i := 2; i < len(coordinates); i++ {
		ratio := getRatio(coordinates[i][0], coordinates[i][1], coordinates[i-1][0], coordinates[i-1][1])
		if ratio != lastRatio {
			return false
		}
	}
	return true
}

func getRatio(x1, y1, x2, y2 int) float64 {
	if x1 == x2 {
		return 0
	} else if y1 == y2 {
		return 1
	}
	return float64(x1-x2) / float64(y1-y2)
}
