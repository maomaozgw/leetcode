package p1455

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	items := strings.Split(sentence, " ")
	for i := range items {
		if strings.HasPrefix(items[i], searchWord) {
			return i + 1
		}
	}
	return -1
}
