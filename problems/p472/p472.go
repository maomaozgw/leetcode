package p472

func findAllConcatenatedWordsInADict(words []string) []string {
	var wordMap = map[string]struct{}{}
	for idx := range words {
		wordMap[words[idx]] = struct{}{}
	}
	var search func(word string) bool
	search = func(word string) bool {
		for i := 0; i < len(word); i++ {
			if _, ok := wordMap[word[:i]]; !ok {
				continue
			}
			if _, ok := wordMap[word[i:]]; ok {
				return true
			}
			if search(word[i:]) {
				return true
			}
		}
		return false
	}
	var result []string
	for idx := range words {
		if search(words[idx]) {
			result = append(result, words[idx])
		}
	}
	return result
}
