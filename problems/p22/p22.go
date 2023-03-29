package p22

func generateParenthesis(n int) []string {
	var result []string

	var dfs func(left, right int, s string)

	dfs = func(left, right int, s string) {
		if len(s) == n*2 {
			result = append(result, s)
			return
		}
		if left < n {
			dfs(left+1, right, s+"(")
		}
		if right < left {
			dfs(left, right+1, s+")")
		}
	}
	dfs(0, 0, "")
	return result
}
