package p503

func nextGreaterElements(nums []int) []int {
	var l = len(nums)
	var result = make([]int, l)
	for i := 0; i < l; i++ {
		result[i] = -1
	}
	for i := 0; i < l; i++ {
		for j := 1; j < l; j++ {
			var k = (i + j) % l
			if nums[k] > nums[i] {
				result[i] = nums[k]
				break
			}
		}
	}
	return result
}
