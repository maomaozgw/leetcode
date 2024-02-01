package p2966

import "sort"

func divideArray(nums []int, k int) [][]int {
	if len(nums)%3 != 0 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, len(nums)/3)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k || nums[i+2]-nums[i+1] > k || nums[i+1]-nums[i] > k {
			return nil
		}
		res[i/3] = []int{nums[i], nums[i+1], nums[i+2]}
	}
	return res
}
