package p985

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	lastSum := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			lastSum += nums[i]
		}
	}
	var result = make([]int, len(queries))
	for i := range queries {
		query := queries[i]
		idx := query[1]
		val := query[0]
		current := nums[idx]
		if current%2 == 0 {
			lastSum -= current
		}
		current += val
		if current%2 == 0 {
			lastSum += current
		}
		nums[idx] = current
		result[i] = lastSum
	}
	return result
}
