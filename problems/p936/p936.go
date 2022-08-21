// https://leetcode.com/problems/stamping-the-sequence/

package p936

import (
	"bytes"
)

func movesToStamp(stamp string, target string) []int {
	var result []int
	l := len(stamp)
	maskLength := 0
	n := len(target)
	maxRound := n * 10
	maskByte := []byte("?")
	replace := bytes.Repeat(maskByte, l)
	tb := []byte(target)
	sb := []byte(stamp)
	tmp := []byte(stamp)
	for {
		hasMatch := false
		for i := l; i > 0; i-- {
			for j := 0; j <= l-i; j++ {
				copy(tmp[:j], replace)
				copy(tmp[j:j+i], sb[j:j+i])
				copy(tmp[j+i:], replace)
				idx := bytes.Index(tb, tmp)
				for idx != -1 {
					hasMatch = true
					maskLength += i
					copy(tb[idx:], replace)
					result = append(result, idx)
					idx = bytes.Index(tb, tmp)
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
