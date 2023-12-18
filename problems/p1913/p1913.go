package p1913

func maxProductDifference(nums []int) int {
	var max0, max1 = 0, 0
	var min0, min1 = 100000000, 100000000
	for i := range nums {
		if nums[i] > max0 {
			max1 = max0
			max0 = nums[i]
		} else if nums[i] > max1 {
			max1 = nums[i]
		}
		if nums[i] < min0 {
			min1 = min0
			min0 = nums[i]
		} else if nums[i] < min1 {
			min1 = nums[i]
		}
	}
	return (max0 * max1) - (min0 * min1)
}
