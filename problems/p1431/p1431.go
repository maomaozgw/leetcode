package p1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxCandyCnt = 0
	for i := range candies {
		maxCandyCnt = max(maxCandyCnt, candies[i])
	}
	var result = make([]bool, len(candies))
	for i := range candies {
		if maxCandyCnt-candies[i] <= extraCandies {
			result[i] = true
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
