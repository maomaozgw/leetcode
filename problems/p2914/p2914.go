package p2914

func minChanges(s string) int {
	var result int
	for i := 0; i < len(s)-1; i += 2 {
		if s[i] != s[i+1] {
			result++
		}
	}
	return result
}
