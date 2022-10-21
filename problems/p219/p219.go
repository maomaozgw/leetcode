// https://leetcode.com/problems/contains-duplicate-ii/

package p219

func containsNearbyDuplicate(nums []int, k int) bool {
	var numLastIdxMap = map[int]int{}
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if lastIdx, ok := numLastIdxMap[n]; ok && (i-lastIdx) <= k {
			return true
		}
		numLastIdxMap[n] = i

	}
	return false
}
