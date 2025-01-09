package p2185

import "strings"

func prefixCount(words []string, pref string) int {
	var result int
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			result++
		}
	}
	return result
}
