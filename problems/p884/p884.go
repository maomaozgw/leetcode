package p884

import (
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	var wordCntMap = map[string]int{}
	var unCommonWordMap = map[string]bool{}
	s1Items := strings.Split(s1, " ")
	for _, w := range s1Items {
		if unCommonWordMap[w] {
			continue
		}
		wordCntMap[w]++
		if wordCntMap[w] > 1 {
			unCommonWordMap[w] = true
			delete(wordCntMap, w)
		}
	}
	s2Items := strings.Split(s2, " ")
	var word2CntMap = map[string]int{}
	for _, w := range s2Items {
		if unCommonWordMap[w] {
			continue
		}
		if wordCntMap[w] > 0 {
			unCommonWordMap[w] = true
			delete(wordCntMap, w)
			continue
		}
		word2CntMap[w]++
		if word2CntMap[w] > 1 {
			unCommonWordMap[w] = true
			delete(word2CntMap, w)
		}
	}
	var result []string
	for w := range wordCntMap {
		result = append(result, w)
	}
	for w := range word2CntMap {
		result = append(result, w)
	}
	return result
}
