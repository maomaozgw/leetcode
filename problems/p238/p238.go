package p238

func productExceptSelf(nums []int) []int {
	var total = 1
	var zeroCnt = 0
	for i := range nums {
		if nums[i] == 0 {
			zeroCnt++
			continue
		}
		total *= nums[i]
	}
	var result = make([]int, len(nums))
	if zeroCnt > 1 {
		return result
	}
	for i := range nums {
		if nums[i] == 0 {
			result[i] = total
			continue
		}
		if zeroCnt == 0 {
			result[i] = total / nums[i]
		}
	}
	return result
}
