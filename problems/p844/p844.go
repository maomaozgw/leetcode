package p844

func backspaceCompare(s string, t string) bool {
	var p1 = len(s) - 1
	var p2 = len(t) - 1

	var cleanFunc = func(s string, idx int) int {
		bCnt := 0
		for ; idx >= 0; idx-- {
			if s[idx] == '#' {
				bCnt++
			} else {
				bCnt--
			}
			if bCnt < 0 {
				return idx
			}
		}
		return idx
	}

	for p1 > -1 || p2 > -1 {
		p1 = cleanFunc(s, p1)
		p2 = cleanFunc(t, p2)
		if p1 > -1 && p2 > -1 && s[p1] != t[p2] {
			return false
		}
		p1--
		p2--
	}
	p1 = cleanFunc(s, p1)
	p2 = cleanFunc(t, p2)
	return p1 == p2
}
