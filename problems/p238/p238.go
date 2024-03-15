package p238

func productExceptSelf(nums []int) []int {
	var prefixProduct, suffixProduct int = 1, 1
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = int(prefixProduct)
		prefixProduct *= num
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * suffixProduct
		suffixProduct *= nums[i]
	}
	return res
}
