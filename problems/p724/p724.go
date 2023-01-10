package p724

func pivotIndex(nums []int) int {
	var sum = 0
	for i := range nums {
		sum += nums[i]
	}
	var left = 0
	for i := 0; i < len(nums); i++ {
		left += nums[i]
		if left == sum {
			return i
		}
		sum -= nums[i]
	}
	return -1
}
