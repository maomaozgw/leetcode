package p1539

func findKthPositive(arr []int, k int) int {
	var left, right = 0, len(arr)
	for left < right {
		var mid = int(uint(left+right) >> 1)
		if arr[mid]-mid > k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left + k
}
