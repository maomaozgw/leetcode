package p1071

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	var m, n = len(str1), len(str2)
	if m < n {
		str1, str2 = str2, str1
		n = m
	}
	for i := 1; i <= n; i++ {
		if n%i != 0 {
			continue
		}
		subStr := str2[:n/i]
		if canMod(str1, subStr) && canMod(str2, subStr) {
			return subStr
		}
	}
	return ""
}

func canMod(str1 string, str2 string) bool {
	m, n := len(str1), len(str2)
	if m%n != 0 {
		return false
	}
	return strings.Repeat(str2, m/n) == str1
}
