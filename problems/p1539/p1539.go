package p1539

func findKthPositive(arr []int, k int) int {
	var j = 0
	var i = 1
	for ; i <= 2000 && k > 0 && j < len(arr); i++ {
		if i == arr[j] {
			j++
			continue
		}
		k--
	}
	return i - 1 + k
}
