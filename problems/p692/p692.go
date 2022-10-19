package p692

import "sort"

func topKFrequent(words []string, k int) []string {
	var freqMap = map[string]int{}
	for i := range words {
		freqMap[words[i]]++
	}
	var result []string
	for s := range freqMap {
		result = append(result, s)
	}
	sort.Slice(result, func(i, j int) bool {
		wi, wj := result[i], result[j]
		fi, fj := freqMap[wi], freqMap[wj]
		if fi > fj {
			return true
		} else if fi < fj {
			return false
		} else {
			return wi < wj
		}
	})

	return result[:k]
}
