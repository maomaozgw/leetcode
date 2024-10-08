package p1963

func minSwaps(s string) int {
	var leftCnt = 0
	var swapCnt = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '[':
			leftCnt++
		case ']':
			if leftCnt == 0 {
				swapCnt++
				leftCnt++
			} else {
				leftCnt--
			}
		}
	}
	return swapCnt
}
