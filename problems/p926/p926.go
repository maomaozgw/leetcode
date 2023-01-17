package p926

func minFlipsMonoIncr(s string) int {
	var oneCnt, result = 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneCnt++
		} else {
			result++
		}
		if result > oneCnt {
			result = oneCnt
		}
	}
	return result
}
