// https: //leetcode.com/problems/combination-sum-iv/

package p377

import (
	"sort"
)

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	resultCache := make([]int, target+1)
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num > i {
				break
			}
			if num == i {
				resultCache[i]++
				break
			}
			remain := i - num
			resultCache[i] += resultCache[remain]

		}
	}
	return resultCache[target]
}
