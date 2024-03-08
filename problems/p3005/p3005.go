package p3005

func maxFrequencyElements(nums []int) int {
	var cntMap = map[int]int{}
	var maxCnt = 0
	for _, num := range nums {
		cntMap[num]++
		if cntMap[num] > maxCnt {
			maxCnt = cntMap[num]
		}
	}
	var result = 0
	for _, cnt := range cntMap {
		if cnt == maxCnt {
			result += cnt
		}
	}
	return result
}
