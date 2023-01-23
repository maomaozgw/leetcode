package p394

func decodeString(s string) string {
	var sb = []byte(s)
	var result []byte
	var i int
	var dfs func()
	dfs = func() {
		var count int
		for i < len(sb) {
			c := s[i]
			i++
			switch c {
			case ']':
				return
			case '[':
				subL := len(result)
				dfs()
				subR := len(result)
				for count--; count != 0; count-- {
					result = append(result, result[subL:subR]...)
				}
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				count = count*10 + int(c-'0')
			default:
				result = append(result, c)
			}
		}
	}
	dfs()
	return string(result)
}
