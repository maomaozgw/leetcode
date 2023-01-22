package p131

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}

	allPartitions := make([][]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		candidate := s[:i+1]
		if palindrome(candidate) {
			remaining := partition(s[i+1:])
			for _, r := range remaining {
				partition := append([]string{candidate}, r...)
				allPartitions = append(allPartitions, partition)
			}
		}
	}

	return allPartitions
}

func palindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return palindrome(s[1 : len(s)-1])
}
