package p11

func maxArea(height []int) int {
	var left, right = 0, len(height) - 1
	var result = 0
	for left < right {
		h := min(height[left], height[right])
		result = max(result, h*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
