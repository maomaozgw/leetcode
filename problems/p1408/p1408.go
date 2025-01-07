package p1408

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var result []string
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}
	return result
}
