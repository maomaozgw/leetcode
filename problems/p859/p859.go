package p859

const (
	startC = 'a'
)

func buddyStrings(s string, goal string) bool {
	var missCnt = 0
	var charCnt = make([]int, 26)
	var fsc, fgc, ssc, sgc byte
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		charCnt[s[i]-startC]++
		if s[i] == goal[i] {
			continue
		}
		missCnt++
		switch missCnt {
		case 1:
			fsc = s[i]
			fgc = goal[i]
		case 2:
			ssc = s[i]
			sgc = goal[i]
		default:
			return false
		}
	}
	if missCnt == 2 {
		return fsc == sgc && fgc == ssc
	} else if missCnt == 1 {
		return false
	}
	for i := range charCnt {
		if charCnt[i] > 1 {
			return true
		}
	}
	return false
}
