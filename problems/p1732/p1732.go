package p1732

func largestAltitude(gain []int) int {
	var maxVal = 0
	var prevVal = 0
	for i := range gain {
		prevVal += gain[i]
		if prevVal > maxVal {
			maxVal = prevVal
		}
	}
	return maxVal
}
