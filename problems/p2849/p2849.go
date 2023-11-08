package p2849

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if sx == fx && sy == fy {
		return t != 1
	}
	xD := abs(sx - fx)
	yD := abs(sy - fy)
	if xD < yD {
		xD = yD
	}
	return xD <= t
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
