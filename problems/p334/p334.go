// https://leetcode.com/problems/increasing-triplet-subsequence/

package p334

const (
	MaxVal = 1<<(31) - 1
)

func increasingTriplet(nums []int) bool {
	var min = MaxVal
	var mid = MaxVal

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n > mid {
			return true
		} else if n > min && n < mid {
			mid = n
		} else if n < min {
			min = n
		}
	}
	return false
}
