package p910

import "sort"

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	var l = len(nums)
	var result = nums[l-1] - nums[0]
	for i := 0; i < l-1; i++ {
		var cMax = max(nums[i]+k, nums[l-1]-k)
		var cMin = min(nums[0]+k, nums[i+1]-k)
		result = min(result, cMax-cMin)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
