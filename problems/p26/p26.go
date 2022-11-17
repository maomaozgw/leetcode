// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package p26

func removeDuplicates(nums []int) int {
	var lastNum = nums[0]
	var writeIdx = 1
	for i := writeIdx; i < len(nums); i++ {
		if nums[i] == lastNum {
			continue
		}
		nums[writeIdx] = nums[i]
		lastNum = nums[i]
		writeIdx++
	}
	return writeIdx
}
