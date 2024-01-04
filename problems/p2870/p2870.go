package p2870

import "sort"

func minOperations(nums []int) int {
	sort.Ints(nums)
	var result int
	for i := 0; i < len(nums); {
		var prev = nums[i]
		var j int
		for j = i; j < len(nums); j++ {
			if nums[j] != prev {
				break
			}
		}
		cnt := j - i
		if cnt == 1 {
			return -1
		}
		result += cnt / 3
		if cnt%3 != 0 {
			result += 1
		}
		i = j
	}
	return result
}
