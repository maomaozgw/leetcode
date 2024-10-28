package p2501

import "sort"

func longestSquareStreak(nums []int) int {
	var (
		result = -1
		numMap = map[int]int{}
		maxNum = -1
	)
	sort.Ints(nums)
	for _, num := range nums {
		numMap[num]++
		maxNum = max(maxNum, num)
	}
	for _, num := range nums {
		count := 1
		for num*num <= maxNum {
			if numMap[num*num] > 0 {
				count++
				num = num * num
			} else {
				break
			}
		}
		if count > 1 {
			result = max(count, result)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
