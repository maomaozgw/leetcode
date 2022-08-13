// https://leetcode.com/problems/substring-with-concatenation-of-all-words/

package p30

func findSubstring(s string, words []string) []int {
	var wordLength = 0
	if len(words) == 0 {
		return nil
	}
	wordLength = len(words[0])
	totalLength := wordLength * len(words)
	var wordMap = map[string]int{}
	for _, w := range words {
		wordMap[w]++
	}
	var result []int
	for i := 0; i <= len(s)-totalLength; i++ {
		subStr := s[i : totalLength+i]
		var hasNoMatch = false
		tmpHashMap := map[string]int{}
		for j := 0; j <= len(subStr)-wordLength; j += wordLength {
			w := subStr[j : wordLength+j]
			if val, ok := wordMap[w]; !ok {
				hasNoMatch = true
				break
			} else {
				tmpHashMap[w]++
				if tmpHashMap[w] > val {
					hasNoMatch = true
					break
				}
			}
		}
		if len(tmpHashMap) != len(wordMap) {
			hasNoMatch = true
		}
		if !hasNoMatch {
			result = append(result, i)
		}
	}
	return result
}
