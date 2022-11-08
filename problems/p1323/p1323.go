// https://leetcode.com/problems/maximum-69-number/

package p1323

func maximum69Number(num int) int {
	var result = num
	var scale = 1
	var maxScale = 0
	for num > 0 {
		current := num % 10
		if current == 6 {
			maxScale = scale
		}
		scale *= 10
		num /= 10
	}

	result += 3 * maxScale
	return result
}
