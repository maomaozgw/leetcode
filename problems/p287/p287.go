package p287

func findDuplicate(nums []int) int {
	var result = map[int]bool{}

	for _, n := range nums {
		if result[n] {
			return n
		}
		result[n] = true
	}
	return -1
}
