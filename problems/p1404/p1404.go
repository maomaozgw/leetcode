package p1404

func numSteps(s string) int {
	l := len(s) - 1
	carry := 0
	count := 0
	for l > 0 {
		t := int(s[l]-'0') + carry
		if t == 0 {
			carry = 0
			count++
		} else if t == 2 {
			carry = 1
			count++
		} else {
			carry = 1
			count += 2
		}
		l--
	}
	if carry == 1 {
		count++
	}
	return count
}
