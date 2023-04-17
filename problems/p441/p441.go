package p441

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	var left, right = 1, n
	for left < right {
		var mid = int(uint(left+right) >> 1)

		var total = (1 + mid) * mid / 2
		if total > n {
			right = mid
		} else if total == n {
			return mid
		} else {
			left = mid + 1
		}

	}
	return left - 1
}
