// https://leetcode.com/problems/valid-anagram/

package p242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sourceMap := map[rune]int{}
	for _, r := range s {
		sourceMap[r] += 1
	}
	for _, r := range t {
		sourceMap[r] -= 1
	}
	for _, c := range sourceMap {
		if c != 0 {
			return false
		}
	}
	return true
}
