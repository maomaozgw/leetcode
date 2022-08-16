// https://leetcode.com/problems/first-unique-character-in-a-string/

package p387

const (
	a = byte('a')
)

func firstUniqChar(s string) int {
	var countSlice = make([]int, 26)
	var indexSlice = make([]int, 26)
	var diffCount = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		idx := c - a
		if countSlice[idx] == 0 {
			indexSlice[idx] = i
		} else if countSlice[idx] == 1 {
			diffCount += 1
			if diffCount == 26 {
				return -1
			}
		}
		countSlice[idx] += 1
	}
	result := len(s)
	for idx := range countSlice {
		if countSlice[idx] != 1 {
			continue
		}
		if indexSlice[idx] < result {
			result = indexSlice[idx]
		}
	}
	if result == len(s) {
		return -1
	}
	return result
}
