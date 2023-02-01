package p1491

func average(salary []int) float64 {
	var min, max = 100000000, 0
	var sum int64
	for _, num := range salary {
		sum += int64(num)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return float64(sum-int64(min)-int64(max)) / float64(len(salary)-2)
}
