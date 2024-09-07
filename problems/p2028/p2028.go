package p2028

func missingRolls(rolls []int, mean int, n int) []int {
	total := (len(rolls) + n) * mean
	for _, v := range rolls {
		total -= v
	}
	if total < n || total > 6*n {
		// out of range
		return nil
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = total / n
		if total%n > 0 {
			result[i]++
			total--
		}
	}
	return result
}
