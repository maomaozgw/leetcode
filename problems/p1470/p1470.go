package p1470

func shuffle(nums []int, n int) []int {
	var result = make([]int, n*2)
	for i := 0; i < n; i++ {
		result[i*2] = nums[i]
		result[i*2+1] = nums[i+n]
	}
	return result
}
