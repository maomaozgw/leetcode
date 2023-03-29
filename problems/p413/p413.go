package p413

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	var l = len(nums)
	var result = 0
	var sameCnt = 0
	var diff int
	var preDiff = nums[1] - nums[0]
	for i := 1; i < l-1; i++ {
		diff = nums[i+1] - nums[i]
		if diff == preDiff {
			sameCnt++
			result += sameCnt
		} else {
			preDiff = diff
			sameCnt = 0
		}
	}
	return result
}
