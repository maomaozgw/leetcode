package p1705

const (
	mod = 1000000007
)

func countHomogenous(s string) int {
	var prevC = byte('0')
	var prevCnt = 0
	var result = 0
	for i := range s {
		if s[i] != prevC {
			prevC = s[i]
			result += ((prevCnt + 1) * prevCnt) / 2
			result = result % mod
			prevCnt = 1
		} else {
			prevCnt++
		}
	}
	result += ((prevCnt + 1) * prevCnt) / 2
	return result % mod
}
