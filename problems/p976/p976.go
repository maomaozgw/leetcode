// https://leetcode.com/problems/largest-perimeter-triangle/

package p976

import "sort"

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	for i := 0; i < len(nums)-2; i++ {
		sum := isTriangle(nums[i:])
		if sum > 0 {
			return sum
		}
	}
	return 0
}

func isTriangle(nums []int) int {
	sum := nums[1] + nums[2]
	if sum > nums[0] {
		sum += nums[0]
		return sum
	}
	return 0
}
