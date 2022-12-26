// https://leetcode.com/problems/sort-characters-by-frequency/

package p451

import (
	"sort"
)

const (
	startC = '0'
)

func frequencySort(s string) string {
	var counter = make([]int, 'z'-startC+1)
	for i := 0; i < len(s); i++ {
		counter[s[i]-startC]++
	}
	var content = []byte(s)
	sort.Slice(content, func(i, j int) bool {
		ci, cj := counter[content[i]-startC], counter[content[j]-startC]
		if ci > cj {
			return true
		} else if ci == cj {
			return content[i] > content[j]
		}
		return false
	})
	return string(content)
}
