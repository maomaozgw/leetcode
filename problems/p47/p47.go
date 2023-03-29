package p47

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var l = len(nums)
	res := [][]int{}
	var frequencies = map[int]int{}
	var backtrack func(list []int)
	for _, n := range nums {
		frequencies[n]++
	}

	backtrack = func(list []int) {
		if len(list) == l {
			tmp := make([]int, l)
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		for k, v := range frequencies {
			if v <= 0 {
				continue
			}
			frequencies[k]--
			backtrack(append(list, k))
			frequencies[k]++
		}
	}

	backtrack([]int{})
	return res
}
