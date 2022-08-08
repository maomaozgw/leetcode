// https://leetcode.com/problems/longest-increasing-subsequence/

package p300

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				result := dp[i] + 1
				if result > dp[j] {
					dp[j] = result
				}
			}
		}
	}
	var largest int
	for _, val := range dp {
		if val > largest {
			largest = val
		}
	}
	return largest + 1
}
