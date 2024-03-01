package p2864

import "strings"

func maximumOddBinaryNumber(s string) string {
	var b = []byte(s)
	var oneCnt = 0
	for _, c := range b {
		if c == '1' {
			oneCnt++
		}
	}
	return strings.Repeat("1", oneCnt-1) + strings.Repeat("0", len(b)-oneCnt) + "1"
}
