package p650

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	var do func(current int, clipboard int) int
	do = func(current int, clipboard int) int {
		if current == n {
			return 0
		}
		if current > n {
			return 10000
		}
		copyAndPaste := 2 + do(current*2, current)
		paste := 1 + do(current+clipboard, clipboard)
		return min(copyAndPaste, paste)
	}
	return do(1, 1) + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
