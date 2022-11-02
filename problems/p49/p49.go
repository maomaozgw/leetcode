package p49

import "sort"

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	var idx = 0
	var indexer = map[string]int{}
	for _, str := range strs {
		ss := sortS(str)
		if foundIdx, ok := indexer[ss]; ok {
			result[foundIdx] = append(result[foundIdx], str)
		} else {
			indexer[ss] = idx
			idx++
			result = append(result, []string{str})
		}
	}
	return result
}

func sortS(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
