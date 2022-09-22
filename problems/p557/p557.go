package p557

func reverseWords(s string) string {
	content := []byte(s)
	var i, j = 0, 0
	for ; i < len(content); i++ {
		if content[i] == ' ' {
			reverse(content, j, i-1)
			j = i + 1
		}
	}
	reverse(content, j, i-1)
	return string(content)
}

func reverse(content []byte, start, end int) {
	if start == end {
		return
	}
	for {
		content[start], content[end] = content[end], content[start]
		if end-start <= 1 {
			return
		}
		start++
		end--
	}
}
