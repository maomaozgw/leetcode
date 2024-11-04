package p3163

import (
	"strconv"
	"strings"
)

func compressedString(word string) string {
	var lastC = word[0]
	var lastCnt = 1
	var sb strings.Builder
	for i := 1; i < len(word); i++ {
		c := word[i]
		if c == lastC && lastCnt < 9 {
			lastCnt++
			continue
		}
		sb.WriteString(strconv.Itoa(lastCnt))
		sb.WriteByte(lastC)
		lastCnt = 1
		lastC = c
	}
	sb.WriteString(strconv.Itoa(lastCnt))
	sb.WriteByte(lastC)
	return sb.String()
}
