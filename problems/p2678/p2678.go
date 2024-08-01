package p2678

func countSeniors(details []string) int {
	var result int
	for i := range details {
		if details[i][11] > '6' {
			result++
		} else if details[i][11] == '6' && details[i][12] > '0' {
			result++
		}
	}
	return result
}
