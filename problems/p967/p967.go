package p967

func numsSameConsecDiff(n int, k int) []int {
	var result []int
	for i := 1; i < 10; i++ {
		if i+k < 10 || i-k >= 0 {
			result = append(result, i)
		}
	}
	for i := 1; i < n; i++ {
		l := len(result)
		var newResult []int
		for j := 0; j < l; j++ {
			last := result[j] % 10
			if last+k < 10 {
				newResult = append(newResult, result[j]*10+last+k)
			}
			if last-k >= 0 && k != 0 {
				newResult = append(newResult, result[j]*10+last-k)
			}
		}
		result = newResult
	}
	return result
}
