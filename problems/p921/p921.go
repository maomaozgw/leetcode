package p921

func minAddToMakeValid(s string) int {
	var leftCnt = 0
	var addCnt = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(':
			leftCnt++
		case ')':
			if leftCnt == 0 {
				addCnt++
			} else {
				leftCnt--
			}
		}
	}
	return leftCnt + addCnt
}
