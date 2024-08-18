package p264

var (
	preData []int
)

func init() {
	var primes = []int{2, 3, 5}
	var indics = []int{0, 0, 0}
	preData = append(preData, 1)
	var dataSet = map[int]bool{
		1: true,
	}
	for len(preData) < 1691 {
		minIdx := 0
		minVal := preData[indics[0]] * primes[0]
		for j := 1; j < 3; j++ {
			if preData[indics[j]]*primes[j] < minVal {
				minIdx = j
				minVal = preData[indics[j]] * primes[j]
			}
		}
		if !dataSet[minVal] {
			dataSet[minVal] = true
			preData = append(preData, minVal)
		}
		indics[minIdx]++
	}
}

func nthUglyNumber(n int) int {
	return preData[n-1]
}
