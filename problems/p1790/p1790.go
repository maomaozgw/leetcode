package p1790

func areAlmostEqual(s1 string, s2 string) bool {
	var l = len(s1)
	var diffCnt = 0
	var srcC = byte(0)
	var dstC = byte(0)

	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			if srcC == 0 {
				srcC = s1[i]
				dstC = s2[i]
			} else {
				if srcC != s2[i] || dstC != s1[i] {
					return false
				}
			}
			diffCnt++
		}
		if diffCnt > 2 {
			return false
		}
	}
	return diffCnt == 2 || diffCnt == 0
}
