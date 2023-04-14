package p367

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	var left, right = 0, num
	for left < right {
		var mid = (left + right) / 2
		var tmp = mid * mid
		if mid == left {
			if tmp == num {
				return true
			} else {
				return false
			}
		}
		if tmp > num {
			right = mid
		} else if tmp == num {
			return true
		} else {
			left = mid
		}
	}
	return false
}
