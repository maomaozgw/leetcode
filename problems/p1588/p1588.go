package p1588

func sumOddLengthSubarrays(arr []int) int {
	var result = 0
	var l = len(arr)
	for i := range arr {
		result += ((i+1)*(l-i) + 1) / 2 * arr[i]
	}
	return result
}
