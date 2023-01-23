package p784

func letterCasePermutation(s string) []string {
	var result []string
	var dfs func(current []byte, idx int)
	dfs = func(current []byte, idx int) {
		if idx == len(s) {
			result = append(result, string(current))
			return
		}
		oldC := s[idx]
		var newC = oldC
		if oldC >= 'a' && oldC <= 'z' {
			newC = oldC - 32
		} else if oldC >= 'A' && oldC <= 'Z' {
			newC = oldC + 32
		}
		dfs(current, idx+1)
		if oldC != newC {
			current[idx] = newC
			dfs(current, idx+1)
			current[idx] = oldC
		}
	}

	dfs([]byte(s), 0)
	return result
}
