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
	var left, right = 0, n
	for left < right {
		var mid = (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
