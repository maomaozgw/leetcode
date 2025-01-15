package p2429

func minimizeXor(num1 int, num2 int) int {
	var num1Ones = countOnes(num1)
	var num2Ones = countOnes(num2)
	if num1Ones == num2Ones {
		return num1
	}
	var result = num1
	for i := 0; i < 32; i++ {
		if num1Ones > num2Ones && (1<<i)&num1 != 0 {
			result ^= 1 << i
			num1Ones--
		}
		if num1Ones < num2Ones && (1<<i)&num1 == 0 {
			result |= 1 << i
			num1Ones++
		}
	}
	return result
}

func countOnes(num int) int {
	var result = 0
	for num > 0 {
		if num&1 == 1 {
			result++
		}
		num >>= 1
	}
	return result
}
