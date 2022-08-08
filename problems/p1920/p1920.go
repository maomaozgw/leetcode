//go:build easy

// https://leetcode.com/problems/build-array-from-permutation/

package p1920

func buildArray(nums []int) []int {
	result := make([]int, len(nums))
	for idx := range nums {
		result[idx] = nums[nums[idx]]
	}
	return result
}
