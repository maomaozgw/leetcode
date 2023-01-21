package p424

func characterReplacement(s string, k int) int {
	var maxLen = 0
	var maxRepeat = byte(0)
	var start = 0
	var end = 0
	var m = make([]int, 128)
	for end < len(s) {
		c := s[end]
		m[c]++
		count := m[c]
		if maxRepeat == 0 || m[maxRepeat] < count {
			maxRepeat = c
		}
		if end-start+1-m[maxRepeat] > k {
			m[s[start]]--
			start++
		}
		if maxLen < end-start+1 {
			maxLen = end - start + 1
		}
		end++
	}
	return maxLen
}
