package p2220

func minBitFlips(start int, goal int) int {
	var result, bita, bitb int
	for start > 0 || goal > 0 {
		bita = start & 1
		bitb = goal & 1
		if bita != bitb {
			result++
		}
		start >>= 1
		goal >>= 1
	}
	return result
}
