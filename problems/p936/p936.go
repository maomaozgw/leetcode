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
	canMatch := func(pos int) int {
		cnt := 0
		for i := 0; i < l; i++ {
			c := tb[i+pos]
			if c == '?' {
				continue
			}
			if c != stamp[i] {
				return -1
			}
			cnt++
		}
		return cnt
	}
	posMatched := make([]bool, n)
	for {
		hasMatch := false
		for i := 0; i < n-l+1; i++ {
			if posMatched[i] {
				continue
			}
			matchCnt := canMatch(i)
			if matchCnt <= 0 {
				continue
			}
			hasMatch = true
			posMatched[i] = true
			copy(tb[i:], replace)
			result = append(result, i)
			maskLength += matchCnt
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
		if !hasMatch {
			break
		}
	}
	return nil
}
