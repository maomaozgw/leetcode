package p3075

import (
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	var result = int64(0)
	for i := 0; i < k; i++ {
		if happiness[i] <= i {
			return result
		}
		result += int64(happiness[i] - i)
	}
	return result
}
