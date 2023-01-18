package p918

func maxSubarraySumCircular(nums []int) int {
	var currentMax, maxSum = 0, -1 << 31
	var currentMin, minSum = 0, 1 << 31
	var total = 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		currentMax += nums[i]
		currentMin += nums[i]
		maxSum = max(maxSum, currentMax)
		if currentMax < 0 {
			currentMax = 0
		}
		minSum = min(minSum, currentMin)
		if currentMin > 0 {
			currentMin = 0
		}

	}
	if total == minSum {
		return maxSum
	}
	return max(maxSum, total-minSum)
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
