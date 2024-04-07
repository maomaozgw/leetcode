package p678

func checkValidString(s string) bool {
	var leftMin, leftMax int
	for i := range s {
		switch s[i] {
		case '(':
			leftMin++
			leftMax++
		case ')':
			leftMin--
			leftMax--
		case '*':
			leftMin--
			leftMax++
		default:
		}
		if leftMax < 0 {
			return false
		}
		if leftMin < 0 {
			leftMin = 0
		}
	}
	return leftMin == 0
}
