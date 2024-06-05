package p260

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	xor &= -xor
	res := []int{0, 0}
	for _, num := range nums {
		if num&xor == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}
