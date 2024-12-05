package p2337

func canChange(start string, target string) bool {
	if start == target {
		return true
	}
	var l, r int
	for i := 0; i < len(start); i++ {
		s := start[i]
		t := target[i]
		if s == 'R' {
			if l > 0 {
				return false
			}
			r++
		}
		if t == 'L' {
			if r > 0 {
				return false
			}
			l++
		}
		if t == 'R' {
			if r == 0 {
				return false
			}
			r--
		}
		if s == 'L' {
			if l == 0 {
				return false
			}
			l--
		}
	}
	return l == 0 && r == 0
}
