package p1318

func minFlips(a int, b int, c int) int {
	var result, bita, bitb, bitc int
	for a > 0 || b > 0 || c > 0 {
		bita = a & 1
		bitb = b & 1
		bitc = c & 1
		if bitc == 0 {
			result += bita + bitb
		} else if bita == 0 && bitb == 0 {
			result++
		}
		a >>= 1
		b >>= 1
		c >>= 1
	}
	return result
}
