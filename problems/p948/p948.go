// https://leetcode.com/problems/bag-of-tokens/

package p948

import (
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})
	var i, j = 0, len(tokens) - 1
	var score = 0
	for i <= j {
		if power >= tokens[i] {
			power -= tokens[i]
			score++
			i++
		} else if score == 0 || i == j {
			return score
		} else {
			score--
			power += tokens[j]
			j--
		}
	}
	return score
}
