package p3254

func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	var result []int
	var increaseCnt = 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			increaseCnt++
		} else {
			increaseCnt = 0
		}
		if increaseCnt >= k-1 {
			result = append(result, nums[i])
		} else {
			result = append(result, -1)
		}
	}
	return result[k-2:]
}
