// https://leetcode.com/problems/word-ladder-ii/

package p126

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var lastIndex = -1
	for idx, str := range wordList {
		if str == endWord {
			lastIndex = idx
			break
		}
	}
	if lastIndex == -1 {
		// Not found
		return [][]string{}
	}

	var stepCache = map[string]int{}
	var proposal [][]string
	proposal = append(proposal, []string{beginWord})
	var nextSlice []string
	nextSlice = append(nextSlice, beginWord)
	stepCache[beginWord] = 0
	for len(nextSlice) > 0 {
		var newNext []string
		for _, w := range nextSlice {
			newNext = append(newNext, foundNext(wordList, stepCache, w)...)
		}
		nextSlice = newNext
	}

	return toList(wordList, stepCache, endWord, beginWord)
}

func toList(wordList []string, cache map[string]int, cw, begin string) [][]string {
	var result [][]string
	if cache[cw] == 1 {
		// 到达终点
		return [][]string{{
			begin, cw,
		}}
	}
	for _, nw := range wordList {
		if val, ok := cache[nw]; !ok {
			continue
		} else if !canWalk(cw, nw) {
			continue
		} else if cache[cw] != val+1 {
			continue
		}
		subResult := toList(wordList, cache, nw, begin)
		for _, item := range subResult {
			item = append(item, cw)
			result = append(result, item)
		}
	}
	return result
}

func foundNext(wordList []string, nextMap map[string]int, fromWord string) []string {
	var result []string
	for _, word := range wordList {
		if word == fromWord {
			continue
		}
		if _, ok := nextMap[word]; ok {
			continue
		}
		if !canWalk(word, fromWord) {
			continue
		}
		nextMap[word] = nextMap[fromWord] + 1
		result = append(result, word)
	}
	return result
}

func canWalk(w1, w2 string) bool {
	count := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return count == 1
}
