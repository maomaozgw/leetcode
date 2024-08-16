package p624

func maxDistance(arrays [][]int) int {
	var smallest, biggest = arrays[0][0], arrays[0][len(arrays[0])-1]
	var result = 0
	for i := 1; i < len(arrays); i++ {
		result = max(result, abs(arrays[i][len(arrays[i])-1]-smallest))
		result = max(result, abs(arrays[i][0]-biggest))
		if arrays[i][0] < smallest {
			smallest = arrays[i][0]
		}
		if arrays[i][len(arrays[i])-1] > biggest {
			biggest = arrays[i][len(arrays[i])-1]
		}
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
