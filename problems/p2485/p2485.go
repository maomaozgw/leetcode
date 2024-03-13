package p2485

func pivotInteger(n int) int {
	return binarySearch(1, n, n)
}

func binarySearch(left, right, n int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	leftSum := (1 + mid) * mid / 2
	rightSum := (mid + n) * (n - mid + 1) / 2
	if leftSum == rightSum {
		return mid
	}
	if leftSum > rightSum {
		return binarySearch(left, mid-1, n)
	}
	return binarySearch(mid+1, right, n)
}
