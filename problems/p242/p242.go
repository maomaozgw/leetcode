// https://leetcode.com/problems/valid-anagram/

package p242

const (
	startC = 'a'
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var cSlice = make([]int, 26)

	for i := range s {
		cSlice[s[i]-startC]++
	}
	for i := range t {
		idx := t[i] - startC
		cSlice[idx]--
		if cSlice[idx] < 0 {
			return false
		}
	}
	return true
}
