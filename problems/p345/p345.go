package p345

var (
	vowelsMap = map[byte]struct {
	}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
)

func reverseVowels(s string) string {
	content := []byte(s)

	var i, j = 0, len(content) - 1
	for {
		for ; i <= j; i++ {
			if _, ok := vowelsMap[content[i]]; ok {
				break
			}
		}
		for ; j >= i; j-- {
			if _, ok := vowelsMap[content[j]]; ok {
				break
			}
		}
		if j > i {
			content[i], content[j] = content[j], content[i]
			j--
			i++
		}
		if j <= i {
			break
		}
	}
	return string(content)
}
