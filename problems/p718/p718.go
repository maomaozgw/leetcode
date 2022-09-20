// https://leetcode.com/problems/maximum-length-of-repeated-subarray/

package p718

func findLength(nums1 []int, nums2 []int) int {
	var dp = make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	var result = 0
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
				result = max(dp[i+1][j+1], result)
			} else {
				dp[i+1][j+1] = 0
			}
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
