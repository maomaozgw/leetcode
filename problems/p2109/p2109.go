package p2109

func addSpaces(s string, spaces []int) string {
	var result = make([]byte, len(s)+len(spaces))
	var j = 0
	var sI = 0
	spaces = append(spaces, -1)
	for i := 0; i < len(s); i++ {
		if i == spaces[sI] {
			sI++
			result[j] = ' '
			j++
		}
		result[j] = s[i]
		j++
	}
	return string(result)
}
