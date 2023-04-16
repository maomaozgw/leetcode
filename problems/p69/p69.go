package p69

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	var left, right = 0, x/2 + 1
	for left < right {
		var mid = (left + right) / 2
		var tmp = mid * mid
		if mid == left {
			if tmp > x {
				return mid - 1
			} else {
				return mid
			}
		}
		if tmp > x {
			right = mid
		} else if tmp == x {
			return mid
		} else {
			left = mid
		}
	}
	return left
}
