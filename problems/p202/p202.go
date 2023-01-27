package p202

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	if n > 0 && n <= 9 {
		if n == 1 || n == 7 {
			return true
		}
		return false
	}
	var sum int
	for n > 0 {
		tmp := n % 10
		sum += tmp * tmp
		n /= 10
	}
	return isHappy(sum)
}
