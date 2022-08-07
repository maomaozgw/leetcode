// https://leetcode.com/problems/running-sum-of-1d-array/

package p1480

func runningSum(nums []int) []int {
	var (
		prev = 0
	)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[prev]
		prev = i
	}
	return nums
}
