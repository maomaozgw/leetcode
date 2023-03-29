package p17

var (
	wordMap = map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var result []string

	var backtrack func(next int, s string)

	backtrack = func(next int, s string) {
		if next == len(digits) {
			result = append(result, s)
			return
		}
		for _, l := range wordMap[digits[next]] {
			backtrack(next+1, s+string(l))
		}
	}
	backtrack(0, "")
	return result
}
