package p278

var (
	bad int
)

func isBadVersion(version int) bool {
	if version >= bad {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	return binarySearch(1, n)
}

func binarySearch(left, right int) int {
	middle := (right + left) / 2
	if middle == left {
		if isBadVersion(middle) {
			return middle
		}
		if isBadVersion(right) {
			return right
		}
		return -1
	}
	if isBadVersion(middle) {
		return binarySearch(left, middle)
	}
	return binarySearch(middle, right)
}
