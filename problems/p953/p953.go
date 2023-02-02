package p953

const (
	start = byte('a')
)

func isAlienSorted(words []string, order string) bool {
	var orderBase = make([]byte, 26)
	for i := 0; i < 26; i++ {
		orderBase[order[i]-start] = byte(i) + start
	}
	for i := 0; i < len(words)-1; i++ {
		if !compare(words[i], words[i+1], orderBase) {
			return false
		}
	}
	return true
}
func compare(s1, s2 string, orderBase []byte) bool {
	l := min(len(s1), len(s2)) + 1
	for i := 0; i < l; i++ {
		c1, c2 := getC(s1, i, orderBase), getC(s2, i, orderBase)
		if c1 > c2 {
			return false
		} else if c1 < c2 {
			return true
		}
	}
	return true
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func getC(s string, index int, mapping []byte) byte {
	if len(s) > index {
		return mapping[s[index]-start]
	}
	return 0
}
