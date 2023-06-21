package p2448

func minCost(nums []int, cost []int) int64 {
	var left, right = 10000000000, 0
	for i := range nums {
		if nums[i] < left {
			left = nums[i]
		}
		if nums[i] > right {
			right = nums[i]
		}
	}

	getCost := func(target int) int {
		var result = 0
		for i := range nums {
			result += abs(nums[i]-target) * cost[i]
		}
		return result
	}
	var result = 0

	for left < right {
		var mid = (left + right) >> 1

		cost := getCost(mid)
		cost1 := getCost(mid + 1)
		if cost > cost1 {
			left = mid + 1
			result = cost1
		} else {
			right = mid
			result = cost
		}
	}
	return int64(result)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
