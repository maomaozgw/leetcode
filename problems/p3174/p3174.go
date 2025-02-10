package p3174

func clearDigits(s string) string {
	var stack []rune
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			// pop last
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
