package p1636

import "sort"

func frequencySort(nums []int) []int {
	var freqMap = map[int]int{}
	for _, num := range nums {
		freqMap[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		f1, f2 := freqMap[nums[i]], freqMap[nums[j]]
		if f1 == f2 {
			return nums[i] > nums[j]
		}
		return f1 < f2
	})
	return nums
}
