package p950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	n := len(deck)
	res := make([]int, n)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	i := 0
	for len(idx) > 0 {
		res[idx[0]] = deck[i]
		idx = idx[1:]
		if len(idx) > 0 {
			idx = append(idx, idx[0])
			idx = idx[1:]
		}
		i++
	}
	return res
}
