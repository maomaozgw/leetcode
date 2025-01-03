package p2270

func waysToSplitArray(nums []int) int {
	var left, right int64
	for _, num := range nums {
		right += int64(num)
	}
	var result int
	for i := 0; i < len(nums)-1; i++ {
		left += int64(nums[i])
		right -= int64(nums[i])
		if left >= right {
			result++
		}
	}
	return result
}
