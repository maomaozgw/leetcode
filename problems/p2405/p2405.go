package p2405

func partitionString(s string) int {
	var result = 1
	var contains = 0
	for idx := range s {
		var c = 1 << (s[idx] - 'a')
		if contains&c == c {
			result++
			contains = c
			continue
		}
		contains |= c
	}
	return result
}
