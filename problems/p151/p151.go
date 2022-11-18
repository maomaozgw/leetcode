// https://leetcode.com/problems/reverse-words-in-a-string/

package p151

import (
	"strings"
)

const (
	space = ' '
)

func reverseWords(s string) string {
	var content = make([]byte, len(s)+1)
	content[0] = ' '
	copy(content[1:], s)
	var builder strings.Builder
	j := len(content) - 1
	for i := len(content) - 2; i >= 0; i-- {
		if content[i] == space && content[i+1] != space {
			if builder.Len() > 0 {
				builder.WriteByte(space)
			}
			builder.Write(content[i+1 : j+1])
		}
		if content[i] != space && content[i+1] == space {
			j = i
		}
	}

	return builder.String()
}
