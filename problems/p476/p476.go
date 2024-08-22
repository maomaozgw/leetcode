package p476

func findComplement(num int) int {
	var result = 0
	var i = 0
	for num > 0 {
		if num&1 == 0 {
			result += 1 << i
		}
		i++
		num >>= 1
	}
	return result
}
