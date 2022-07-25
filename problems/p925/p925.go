package p925

func isLongPressedName(name string, typed string) bool {
	var (
		tInx  = 0
		tLen  = len(typed)
		nInx  = 0
		nCnt  = 0
		lastN uint8
	)
	for ; nInx < len(name)-1; nInx++ {
		nCnt++
		lastN = name[nInx]
		if name[nInx+1] == name[nInx] {
			continue
		}

		for ; tInx < tLen; tInx++ {
			tc := typed[tInx]
			if tc != lastN {
				if nCnt <= 0 {
					break
				}
				return false
			}
			nCnt--
		}
		if nCnt > 0 {
			return false
		}
		nCnt = 0
	}
	lastN = name[nInx]
	nCnt = 1
	for ; tInx < tLen; tInx++ {
		tc := typed[tInx]
		if tc != lastN {
			return false
		}
		nCnt -= 1
	}
	return nCnt <= 0
}
