package p633

func judgeSquareSum(c int) bool {
	var left, right = 0, square(c)
	for left < right {
		var sum = left*left + right*right
		// var mid = (left + right) >> 1
		// if left*left+mid*mid > c {
		// 	right = mid
		// 	continue
		// } else if right*right+mid*mid < c {
		// 	left = mid + 1
		// 	continue
		// }
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return left*left*2 == c
}

func square(num int) int {
	var left, right = 0, num
	for left < right {
		var mid = int((uint(left) + uint(right)) >> 1)
		var tmp = mid * mid
		if tmp > num {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
