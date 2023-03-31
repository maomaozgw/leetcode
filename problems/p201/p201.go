package p201

func rangeBitwiseAnd(left int, right int) int {
	var diff = right - left
	var cnt = 0
	for diff > 0 {
		diff >>= 1
		cnt++
	}
	return ((left & right) >> cnt) << cnt
}
