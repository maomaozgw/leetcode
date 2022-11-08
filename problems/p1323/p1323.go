// https://leetcode.com/problems/maximum-69-number/

package p1323

func maximum69Number(num int) int {
	var bytes []byte

	for num > 0 {
		bytes = append(bytes, uint8(num%10))
		num /= 10
	}

	var transferred = false

	for i := len(bytes) - 1; i > -1; i-- {
		if !transferred && bytes[i] == 6 {
			bytes[i] = 9
			transferred = true
		}
		num *= 10
		num += int(bytes[i])
	}
	return num
}
