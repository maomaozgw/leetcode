package p3016

import "sort"

func minimumPushes(word string) int {
	if len(word) <= 8 {
		return len(word)
	}
	var counter = make([]int, 26)
	for _, c := range word {
		counter[c-'a']++
	}
	var res int
	sort.Slice(counter, func(i, j int) bool {
		return counter[i] > counter[j]
	})
	for idx, v := range counter {
		if v == 0 {
			break
		}
		res += ((idx / 8) + 1) * v
	}
	return res
}
