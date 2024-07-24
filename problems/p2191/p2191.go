package p2191

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	var cache = make(map[int]int)
	var idxMap = make(map[int]int)
	var getMapping func(num int) int
	getMapping = func(num int) int {
		if num < 10 {
			return mapping[num]
		}
		if target, ok := cache[num]; ok {
			return target
		}
		result := getMapping(num/10)*10 + mapping[num%10]
		cache[num] = result
		return result
	}
	for i, num := range nums {
		idxMap[num] = i
	}
	sort.Slice(nums, func(i, j int) bool {
		m1, m2 := getMapping(nums[i]), getMapping(nums[j])
		if m1 == m2 {
			return idxMap[nums[i]] < idxMap[nums[j]]
		}
		return m1 < m2
	})
	return nums
}
