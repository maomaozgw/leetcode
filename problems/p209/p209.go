package p209

func minSubArrayLen(target int, nums []int) int {
	var l = len(nums)
	var result = l + 1
	var val = 0
	var left = 0
	for i := 0; i < l; {
		if val >= target {
			val -= nums[left]
			left++
			result = min(result, i-left+1)
			continue
		}
		val += nums[i]
		i++
	}
	for val >= target {
		val -= nums[left]
		left++
		result = min(result, l-left+1)
	}
	if result > l {
		return 0
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
