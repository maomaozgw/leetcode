// https://leetcode.com/problems/stamping-the-sequence/

package p936

import (
	"strings"
)

func movesToStamp(stamp string, target string) []int {
	var result []int
	l := len(stamp)
	maskLength := 0
	n := len(target)
	maxRound := n * 10
	replace := strings.Repeat("?", l)
	for {
		hasMatch := false
		for i := l; i > 0; i-- {
			for j := 0; j <= l-i; j++ {
				tmp := strings.Repeat("?", j) + stamp[j:j+i] + strings.Repeat("?", l-i-j)
				idx := strings.Index(target, tmp)
				for idx != -1 {
					hasMatch = true
					maskLength += i
					target = strings.Replace(target, tmp, replace, 1)
					result = append(result, idx)
					idx = strings.Index(target, tmp)
				}
				if maskLength == n {
					for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
						result[i], result[j] = result[j], result[i]
					}
					return result
				}
				if len(result) > maxRound {
					return nil
				}
			}
		}
		if !hasMatch {
			break
		}
	}
	return nil
}
