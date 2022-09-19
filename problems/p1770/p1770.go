// https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/

package p1770

func maximumScore(nums []int, multipliers []int) int {
	var dp = make([]int, len(multipliers)+1)
	for i := len(multipliers) - 1; i >= 0; i-- {
		var row = make([]int, len(multipliers)+1)
		copy(row, dp)
		for l := i; l >= 0; l-- {
			dp[l] = max(
				multipliers[i]*nums[l]+row[l+1],
				multipliers[i]*nums[len(nums)-1-i+l]+row[l],
			)
		}
	}
	return dp[0]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
