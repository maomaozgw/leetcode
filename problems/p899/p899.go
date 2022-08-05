// https://leetcode.com/problems/orderly-queue/

package p899

import (
	"sort"
)

func orderlyQueue(s string, k int) string {
	if k >= 2 {
		content := []byte(s)
		sort.Slice(content, func(i, j int) bool {
			return content[i] < content[j]
		})
		return string(content)
	}
	var smallest = s
	var current = s
	for i := 0; i < len(s); i++ {
		current = current[1:] + current[:1]
		if current < smallest {
			smallest = current
		}
	}
	return smallest
}
