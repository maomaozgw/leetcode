package p2256

func minimumAverageDifference(nums []int) int {
	sum := 0
	var l = len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	var minVal = 1<<(31) - 1
	var minIndex = -1
	var leftPart = 0
	var rightPart = sum
	for i := 0; i < l; i++ {
		leftPart += nums[i]
		rightPart -= nums[i]
		val := leftPart / (i + 1)
		if i < l-1 {
			val = abs(val - rightPart/(l-1-i))
		}
		if val < minVal {
			minVal = val
			minIndex = i
			if val == 0 {
				return i
			}
		}
	}
	return minIndex
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
