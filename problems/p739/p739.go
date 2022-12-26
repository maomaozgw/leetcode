package p739

const (
	maxNum = 100
)

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	var records = make([]int, 101)
	for i := 0; i < len(records); i++ {
		records[i] = -1
	}
	var result = make([]int, l)
	for i := l - 1; i >= 0; i-- {
		records[temperatures[i]] = i
		if temperatures[i] >= maxNum {
			result[i] = 0
			continue
		}
		minIdx := l
		for k := temperatures[i] + 1; k <= maxNum; k++ {
			if records[k] > -1 {
				if records[k] < minIdx {
					minIdx = records[k]
				}
			}
		}
		if minIdx == l {
			result[i] = 0
		} else {
			result[i] = minIdx - i
		}
	}
	return result
}
