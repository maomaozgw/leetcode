package p392

func isSubsequence(s string, t string) bool {
	var i, j = 0, 0
	if len(s) == 0 {
		return true
	}
	for ; i < len(s); i++ {
		var found = false
		for ; j < len(t); j++ {
			if s[i] == t[j] {
				j++
				found = true
				break
			}
		}
		if i == len(s)-1 && found {
			return true
		}
		if j == len(t) {
			return false
		}
	}
	return false
}
