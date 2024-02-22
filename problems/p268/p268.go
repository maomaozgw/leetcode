package p268

func missingNumber(nums []int) int {
	var result = 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
		result ^= i
	}
	result ^= len(nums)
	return result
}
