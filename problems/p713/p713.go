package p713

func numSubarrayProductLessThanK(nums []int, k int) int {
	var left, right = 0, 0
	var prod = 1
	var count = 0
	var l = len(nums)

	for right < l {
		prod *= nums[right]
		for prod >= k && left <= right {
			prod /= nums[left]
			left++
		}
		count += right - left + 1
		right++
	}
	return count
}
