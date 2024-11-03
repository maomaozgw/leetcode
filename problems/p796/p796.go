package p796

func rotateString(s string, goal string) bool {
	if s == goal {
		return true
	}

	for i := 1; i < len(s); i++ {
		if s[i:]+s[:i] == goal {
			return true
		}
	}

	return false
}
