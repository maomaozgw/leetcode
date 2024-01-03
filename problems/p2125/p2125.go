package p2125

import "strings"

func numberOfBeams(bank []string) int {
	var result int
	var prevC int
	for i := 0; i < len(bank); i++ {
		oCnt := strings.Count(bank[i], "1")
		if oCnt == 0 {
			continue
		}
		if prevC == 0 {
			prevC = oCnt
			continue
		}
		result += oCnt * prevC
		prevC = oCnt
	}
	return result
}
