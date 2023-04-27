package p1760

func minimumSize(nums []int, maxOperations int) int {
	var left = 1
	var right = 0
	var result = 0
	for i := range nums {
		right = max(right, nums[i])
	}
	for left <= right {
		var operations = 0
		var mid = (right-left)>>1 + left
		for i := range nums {
			if nums[i]%mid == 0 {
				operations += nums[i]/mid - 1
			} else {
				operations += nums[i] / mid
			}
		}
		if operations > maxOperations {
			left = mid + 1
		} else {
			right = mid - 1
			result = mid
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
