package p1750

func minimumLength(s string) int {
	var left, right = 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			break
		}
		var c = s[left]
		for left <= right && s[left] == c {
			left++
		}
		for left <= right && s[right] == c {
			right--
		}
	}
	return right - left + 1
}
