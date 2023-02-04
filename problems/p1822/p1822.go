package p1822

func arraySign(nums []int) int {
	var cnt = 0
	for i := range nums {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			cnt++
		}
	}
	if cnt%2 == 0 {
		return 1
	}
	return -1
}
