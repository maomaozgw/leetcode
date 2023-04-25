package p1894

func chalkReplacer(chalk []int, k int) int {
	var sum = 0
	for i := range chalk {
		sum += chalk[i]
		k -= chalk[i]
		if k < 0 {
			return i
		}
	}
	k %= sum
	if k == 0 {
		return 0
	}
	for i := range chalk {
		k -= chalk[i]
		if k < 0 {
			return i
		}
	}
	return len(chalk) - 1
}
