// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/

package p2131

func longestPalindrome(words []string) int {
	var (
		palindromeCount = 0
	)
	var wMap = map[string]int{}
	var sMap = map[string]int{}
	for _, w := range words {
		rw := reverse(w)
		if w == rw {
			sMap[w]++
			continue
		}
		if wMap[rw] > 0 {
			palindromeCount += 2
			wMap[rw]--
			continue
		}
		wMap[w]++

	}
	var total = 0
	total += palindromeCount
	var hasSingle = false
	for _, count := range sMap {
		mod := count % 2
		if mod == 1 {
			hasSingle = true
		}
		total += count - mod
	}
	if hasSingle {
		return (total + 1) * 2
	}
	return total * 2
}

func reverse(s string) string {
	bs := []byte(s)
	bs[0], bs[1] = bs[1], bs[0]
	return string(bs)
}
