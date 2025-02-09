package p2364

func countBadPairs(nums []int) int64 {
	var result int64
	var counter = map[int]int{}
	for i, num := range nums {
		result += int64(i - counter[num-i])
		counter[num-i]++
	}
	return result
}
