package p16

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var result = 0
	var minGap = 20000
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			gap := sum - target
			if abs(gap) < minGap {
				minGap = abs(gap)
				result = sum
			}
			if gap == 0 {
				return target
			} else if gap > 0 {
				k--
			} else if gap < 0 {
				j++
			}
		}
	}
	return result
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
