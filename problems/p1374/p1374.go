// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/

package p1374

func generateTheString(n int) string {
	if n == 1 {
		return "a"
	}
	var content = make([]byte, n)
	if n%2 == 0 {
		content[n-1] = 'b'
		n--
	}
	for i := 0; i < n; i++ {
		content[i] = 'a'
	}
	return string(content)
}
