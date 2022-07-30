// https://leetcode.com/problems/word-subsets/

package p916

func wordSubsets(words1 []string, words2 []string) []string {
	var result []string
	bSet, diffCount := getSet(words2)
	for _, w1 := range words1 {
		tmpSet := make([]int, 26)
		copy(tmpSet, bSet)
		if isSubset(w1, tmpSet, diffCount) {
			result = append(result, w1)
		}

	}
	return result
}

func getSet(bList []string) ([]int, int) {
	var bSet = make([]int, 26)
	for _, b := range bList {
		tmpSet := make([]int, 26)
		for idx := 0; idx < len(b); idx++ {
			tmpSet[b[idx]-'a'] += 1
		}
		for idx := 0; idx < 26; idx++ {
			if tmpSet[idx] > bSet[idx] {
				bSet[idx] = tmpSet[idx]
			}
		}
	}
	diffCount := 0
	for _, v := range bSet {
		if v > 0 {
			diffCount++
		}
	}
	return bSet, diffCount
}

func isSubset(a string, bSet []int, diffCount int) bool {
	if len(a) < diffCount {
		return false
	}
	for idx := 0; idx < len(a); idx++ {
		c := a[idx] - 'a'
		bSet[c]--
	}
	for _, v := range bSet {
		if v > 0 {
			return false
		}
	}
	return true
}
