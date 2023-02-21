package p1768

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var p1, p2 = 0, 0
	var result strings.Builder
	for p1 < len(word1) && p2 < len(word2) {
		result.WriteByte(word1[p1])
		result.WriteByte(word2[p2])
		p1++
		p2++
	}
	for p1 < len(word1) {
		result.WriteByte(word1[p1])
		p1++
	}
	for p2 < len(word2) {
		result.WriteByte(word2[p2])
		p2++
	}
	return result.String()
}
