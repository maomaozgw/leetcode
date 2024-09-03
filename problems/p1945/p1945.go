package p1945

func getLucky(s string, k int) int {
	var result int
	for _, c := range s {
		v := int(c - 'a' + 1)
		for v > 0 {
			result += v % 10
			v /= 10
		}
	}
	for i := 0; i < k-1; i++ {
		var current int
		for result > 0 {
			current += result % 10
			result /= 10
		}
		result = current
	}
	return result
}
