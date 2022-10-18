// https://leetcode.com/problems/count-and-say/

package p38

import (
	"strconv"
	"strings"
)

var saySlice []string

func init() {
	saySlice = make([]string, 30)
	saySlice[0] = "1"
	for i := 1; i < 30; i++ {
		saySlice[i] = say(saySlice[i-1])
	}
}

func say(origin string) string {
	sb := strings.Builder{}
	var lastC = origin[0]
	var lastCnt = 1
	for i := 1; i < len(origin); i++ {
		c := origin[i]
		if c == lastC {
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

func countAndSay(n int) string {
	return saySlice[n-1]
}
