package p40

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var backtrack func(idx, t int, list []int)

	backtrack = func(idx, t int, list []int) {
		if t == 0 {
			tmp := make([]int, len(list))
			copy(tmp, list)
			result = append(result, tmp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > t {
				break
			}
			backtrack(i+1, t-candidates[i], append(list, candidates[i]))
		}
	}
	backtrack(0, target, []int{})
	return result
}
