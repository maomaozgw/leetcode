package p2962

func countSubarrays(nums []int, k int) int64 {
	var maxNum int
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	var result int64
	var cnt int
	var left = 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == maxNum {
			cnt++
		}
		for cnt >= k {
			if nums[left] == maxNum {
				cnt--
			}
			left++
		}
		result += int64(left)
	}
	return result
}
