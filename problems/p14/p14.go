package p14

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return ""
	}

	var c uint8
	l := len(strs)
	for index := 0; ; index++ {
		if index == len(strs[0]) {
			return result
		}
		c = strs[0][index]

		for j := 1; j < l; j++ {
			if index == len(strs[j]) {
				return result
			}
			if strs[j][index] != c {
				return result
			}
		}
		result += string(c)
	}
}
