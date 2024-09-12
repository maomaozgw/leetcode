package p1684

func countConsistentStrings(allowed string, words []string) int {
	var allowedMap = make([]bool, 26)
	for _, ch := range allowed {
		allowedMap[ch-'a'] = true
	}
	var result int
	for _, word := range words {
		consistent := true
		for _, ch := range word {
			if !allowedMap[ch-'a'] {
				consistent = false
				break
			}
		}
		if consistent {
			result++
		}
	}
	return result
}
