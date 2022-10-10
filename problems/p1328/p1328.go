// https://leetcode.com/problems/break-a-palindrome/

package p1328

func breakPalindrome(palindrome string) string {
	var l = len(palindrome)
	if l <= 1 {
		return ""
	}
	var n = l / 2
	var b = []byte(palindrome)
	for i := 0; i < n; i++ {
		if b[i] != 'a' {
			b[i] = 'a'
			return string(b)
		}
	}
	b[l-1] = 'b'
	return string(b)
}
