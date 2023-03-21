package p2348

func zeroFilledSubarray(nums []int) int64 {
	var result = int64(0)
	var zeroCnt = int64(0)
	for i := range nums {
		if nums[i] == 0 {
			zeroCnt++
			continue
		}
		if zeroCnt == 1 {
			result++

		} else if zeroCnt > 1 {
			result += (1 + zeroCnt) * zeroCnt / 2
		}
		zeroCnt = 0
	}
	if zeroCnt == 1 {
		result++
	} else if zeroCnt > 1 {
		result += (1 + zeroCnt) * zeroCnt / 2
	}
	return int64(result)
}
