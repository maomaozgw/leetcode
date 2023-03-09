package p1208

func equalSubstring(s string, t string, maxCost int) int {
	var nums = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		nums[i] = abs(int(s[i]) - int(t[i]))
	}
	var left, right = 0, 0
	var total = 0
	var result = 0
	for right < len(s) {
		total += nums[right]
		for total > maxCost {
			total -= nums[left]
			left++
		}
		right++
		result = max(result, right-left)
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
