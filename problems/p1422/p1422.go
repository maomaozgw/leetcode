package p1422

func maxScore(s string) int {
	var left = 0
	var right = 0
	for i := range s {
		if s[i] == '0' {
			continue
		}
		right++
	}
	var result = 0
	for i := 0; i < len(s)-1; i++ {

		if s[i] == '1' {
			right--
		} else {
			left++
		}
		result = max(result, left+right)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
