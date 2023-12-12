package p1287

func findSpecialInteger(arr []int) int {
	var minCnt = len(arr) / 4
	var prevCnt = 0
	var prevN = -1
	for i := 0; i < len(arr); i++ {
		if prevCnt > minCnt {
			return prevN
		}
		if arr[i] == prevN {
			prevCnt++
			continue
		}
		prevN = arr[i]
		prevCnt = 1
	}
	return prevN
}
