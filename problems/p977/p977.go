package p977

func sortedSquares(nums []int) []int {
	var left, right = 0, len(nums) - 1
	var result = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			result[i] = nums[right] * nums[right]
			right--
		} else {
			result[i] = nums[left] * nums[left]
			left++
		}
	}
	return result
}
