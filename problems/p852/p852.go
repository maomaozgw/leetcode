package p852

func peakIndexInMountainArray(arr []int) int {
	var left = 0
	var right = len(arr) - 1
	for left < right {
		var mid = (left + right) / 2
		if mid == left {
			return left
		}
		if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			left = mid
		} else if arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1] {
			right = mid
		} else {
			return mid
		}
	}
	return left
}
