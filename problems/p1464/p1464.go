package p1464

func maxProduct(nums []int) int {
	var max1, max2 int
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	return (max1 - 1) * (max2 - 1)
}
