// https://leetcode.com/problems/roman-to-integer/

package p13

var (
	valMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func romanToInt(s string) int {
	var result = 0
	var currentMax = 0
	for l := len(s) - 1; l >= 0; l-- {
		val := valMap[s[l]]
		if val >= currentMax {
			currentMax = val
			result += val
		} else {
			result -= val
		}
	}
	return result
}
