package p409

const start = 'A'

func longestPalindrome(s string) int {
	var runes = make([]int, int('z'-'A')+1)
	for i := 0; i < len(s); i++ {
		runes[s[i]-start]++
	}
	var hasSingle = false
	var sameCount = 0
	for i := range runes {
		if runes[i]%2 == 1 {
			hasSingle = true
		}
		sameCount += runes[i] / 2
	}
	var result = sameCount * 2
	if hasSingle {
		result += 1
	}
	return result
}
