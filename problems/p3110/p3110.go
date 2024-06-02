package p3110

func scoreOfString(s string) int {
	var result int
	for i := 1; i < len(s); i++ {
		result += abs(s[i], s[i-1])
	}
	return result
}

func abs(a, b byte) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}
