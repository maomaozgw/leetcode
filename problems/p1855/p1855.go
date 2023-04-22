package p1855

func maxDistance(nums1 []int, nums2 []int) int {
	var result = 0
	var l2 = len(nums2)
	for i := 0; i < len(nums1); i++ {
		if l2-i < result {
			break
		}
		var left, right = i, l2
		for left < right {
			var mid = (left + right) >> 1
			if nums2[mid] < nums1[i] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		result = max(result, left-i-1)
	}
	return result
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
