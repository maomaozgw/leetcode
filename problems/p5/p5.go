package p5

func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	subStrL := l
	for subStrL > 0 {
		for i := 0; i <= l-subStrL; i++ {
			if isPalindrome(s[i : i+subStrL]) {
				return s[i : i+subStrL]
			}
		}
		subStrL--
	}
	return s[:1]
}

func isPalindrome(a string) bool {
	var l = len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[l-i] {
			return false
		}
	}
	return true
}
