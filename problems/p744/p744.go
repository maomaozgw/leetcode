package p744

func nextGreatestLetter(letters []byte, target byte) byte {
	var left, right = 0, len(letters)

	for left < right {
		var mid = (left + right) / 2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == len(letters) {
		return letters[0]
	}
	return letters[left]
}
