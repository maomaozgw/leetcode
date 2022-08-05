// https://leetcode.com/problems/concatenation-of-array/

package p1929

func getConcatenation(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	var result = make([]int, len(nums)*2)
	copy(result, nums)
	copy(result[len(nums):], nums)
	return result
}
