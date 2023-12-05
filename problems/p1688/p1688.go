package p1688

func numberOfMatches(n int) int {
	var result = 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			result += n
		} else {
			n = n/2 + 1
			result += n - 1
		}
	}
	return result
}
