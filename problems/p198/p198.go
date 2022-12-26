// https://leetcode.com/problems/house-robber/description/

package p198

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	l := len(nums)
	var max1, max2 = nums[l-2], nums[l-1]
	for i := l - 3; i > -1; i-- {
		current := nums[i]
		currentMax := max(current+max2, current+max1-nums[i+1])
		max2, max1 = max1, currentMax
	}
	return max(max1, max2)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
