package p1523

func countOdds(low int, high int) int {
	result := (high - low) / 2
	if low%2 == 1 || high%2 == 1 {
		result++
	}
	return result
}
