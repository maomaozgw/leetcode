package p567

const (
	start = 'a'
)

func checkInclusion(s1 string, s2 string) bool {
	var charSet = make([]int, 26)
	for i := 0; i < len(s1); i++ {
		charSet[s1[i]-start]++
	}
	var tmpSet = make([]int, 26)
	var j = 0
	for i := 0; i < len(s2); i++ {
		c := s2[i] - start
		tmpSet[c]++
		var cnt = i - j + 1
		if cnt == len(s1) {
			if equals(charSet, tmpSet) {
				return true
			}
		}
		if cnt >= len(s1) {
			tmpSet[s2[j]-start]--
			j++
		}
	}
	return false
}

func equals(set1, set2 []int) bool {
	for i := 0; i < len(set1); i++ {
		if set1[i] != set2[i] {
			return false
		}
	}
	return true
}
