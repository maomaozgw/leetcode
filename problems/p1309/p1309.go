package p1309

const (
	singleOffset = 'a' - '1'
	oneOffset    = 'j' - '0'
	twoOffset    = 'j' + 10 - '0'
)

func freqAlphabets(s string) string {
	var result []byte
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == '#' {
			i -= 2
			if s[i] == '1' {
				result = append(result, s[i+1]+oneOffset)
			} else {
				result = append(result, s[i+1]+twoOffset)
			}
			continue
		}
		result = append(result, s[i]+singleOffset)
	}
	var l = len(result)
	for i := 0; i < l/2; i++ {
		result[i], result[l-i-1] = result[l-i-1], result[i]
	}
	return string(result)
}
