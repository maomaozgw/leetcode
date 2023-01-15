package p205

func isIsomorphic(s string, t string) bool {
	l := len(s)
	var prevS, prevT = byte(0), byte(0)
	var cMap = make([]byte, 128)
	var bMap = make([]byte, 128)
	for i := 0; i < l; i++ {
		if cMap[s[i]] == 0 {
			cMap[s[i]] = t[i]
		} else if cMap[s[i]] != t[i] {
			return false
		}
		if bMap[t[i]] == 0 {
			bMap[t[i]] = s[i]
		} else if bMap[t[i]] != s[i] {
			return false
		}
		if (s[i] == prevS) == (t[i] == prevT) {
			prevS, prevT = s[i], t[i]
			cMap[s[i]] = t[i]
			continue
		}
		return false
	}
	return true
}
