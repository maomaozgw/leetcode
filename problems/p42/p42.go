// https://leetcode.com/problems/trapping-rain-water/

package p42

func trap(height []int) int {
	l := len(height)
	var leftMax = make([]int, l-1)
	var rightMax = make([]int, l-1)
	var preMax = 0
	for i := 0; i < l-1; i++ {
		leftMax[i] = max(preMax, height[i])
		preMax = leftMax[i]
	}
	preMax = 0
	for i := l - 1; i > 0; i-- {
		rightMax[i-1] = max(preMax, height[i])
		preMax = rightMax[i-1]
	}
	var result = 0
	for i := 1; i < l-1; i++ {
		current := min(leftMax[i-1], rightMax[i]) - height[i]
		if current > 0 {
			result += current
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
