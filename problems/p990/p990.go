// https://leetcode.com/problems/satisfiability-of-equality-equations/

package p990

import "fmt"

var (
	cValues []int
)

func init() {
	for i := 0; i < 26; i++ {
		cValues = append(cValues, 1<<i)
	}
}

func equationsPossible(equations []string) bool {
	var equals = make([]int, 26)
	copy(equals, cValues)
	var unEquals = make([]int, 26)
	for idx := range equations {
		item := equations[idx]
		c1 := item[0] - 'a'
		c2 := item[3] - 'a'
		if item[1] == '!' {
			unEquals[c1] |= cValues[c2]
			unEquals[c2] |= cValues[c1]
		} else {
			v1 := cValues[c1]
			v2 := cValues[c2]
			equals[c1] |= v2
			equals[c2] |= v1
			for eIdx := range equals {
				if equals[eIdx]&v1 != 0 {
					equals[eIdx] |= equals[c1]
				}
				if equals[eIdx]&v2 != 0 {
					equals[eIdx] |= equals[c2]
				}
			}
		}
	}
	fmt.Println(equals, unEquals)
	for idx := range equals {
		if equals[idx]&unEquals[idx] != 0 {
			return false
		}
	}
	return true
}
