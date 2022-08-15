// https://leetcode.com/problems/integer-to-roman/

package p12

var valueSlice = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var valueStr = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	var result string
	for i := 0; i < len(valueSlice)-1; {
		val := valueSlice[i]
		if num >= val {
			num -= val
			result += valueStr[i]
		} else {
			i++
		}
	}
	for i := 0; i < num; i++ {
		result += "I"
	}
	return result
}
