package p77

func combine(n int, k int) [][]int {
	var result [][]int
	var combineFunc func(start int, current []int)
	combineFunc = func(start int, current []int) {
		if len(current) == k {
			var tmp []int
			tmp = append(tmp, current...)
			result = append(result, tmp)
			return
		}
		for ; start <= n; start++ {
			current = append(current, start)
			combineFunc(start+1, current)
			current = current[:len(current)-1]
		}
	}
	combineFunc(1, []int{})
	return result
}
