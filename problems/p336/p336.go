// https://leetcode.com/problems/palindrome-pairs/

package p336

import (
	"strings"
)

func reverse(s string) string {
	r := strings.Builder{}
	for i := len(s) - 1; i > -1; i-- {
		r.WriteByte(s[i])
	}
	return r.String()
}

func palindromePairs(words []string) [][]int {
	var backWords = map[string]int{}
	var emptyIdx = -1
	for i := range words {
		if words[i] == "" {
			emptyIdx = i
			continue
		}
		backWords[reverse(words[i])] = i
	}

	var result [][]int
	for i := range words {
		w := words[i]
		if bi, ok := backWords[w]; ok && bi == i && emptyIdx > -1 && len(w) > 0 {
			result = append(result, []int{i, emptyIdx})
			result = append(result, []int{emptyIdx, i})
		}

		for j := 0; j < len(w); j++ {
			right := w[j:]
			left := w[:j]
			if bi, ok := backWords[left]; ok && bi != i {
				if right == reverse(right) {
					result = append(result, []int{i, bi})
				}
			}
			if bi, ok := backWords[right]; ok && bi != i {
				if left == reverse(left) {
					result = append(result, []int{bi, i})
				}
			}
		}
	}
	return result
}

func palindromePairsTL(words []string) [][]int {
	var result [][]int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isValid(words[i], words[j]) {
				result = append(result, []int{i, j})
			}
			if isValid(words[j], words[i]) {
				result = append(result, []int{j, i})
			}
		}
	}
	return result
}

func isValid(part1, part2 string) bool {
	l1 := len(part1)
	l2 := len(part2)
	lt := l1 + l2 - 1
	for i := 0; i <= lt/2; i++ {
		j := lt - i
		if i >= l1 {
			if part2[i-l1] != part2[j-l1] {
				return false
			}
		} else if j < l1 {
			if part1[i] != part1[j] {
				return false
			}
		} else {
			if part1[i] != part2[j-l1] {
				return false
			}
		}
	}
	return true
}
