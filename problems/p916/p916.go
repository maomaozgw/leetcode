// https://leetcode.com/problems/word-subsets/

package p916

func wordSubsets(words1 []string, words2 []string) []string {
	var result []string
	var bSet = map[uint8]int{}
	for _, b := range words2 {
		tmpSet := getSet(b)
		for c, v := range tmpSet {
			if v2 := bSet[c]; v > v2 {
				bSet[c] = v
			}
		}
	}
	for _, w1 := range words1 {
		if len(w1) < len(bSet) {
			continue
		}
		if isSubset(w1, bSet) {
			result = append(result, w1)
		}

	}
	return result
}

func getSet(str string) map[uint8]int {
	tmpSet := map[uint8]int{}
	for idx := 0; idx < len(str); idx++ {
		tmpSet[str[idx]-'a'] += 1
	}
	return tmpSet
}

func isSubset(a string, bSet map[uint8]int) bool {
	aSet := getSet(a)
	for c, v := range bSet {
		if av := aSet[c]; v > av {
			return false
		}
	}
	return true
}
