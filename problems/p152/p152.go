package p152

func maxProduct(nums []int) int {
	var maxSoFar int = nums[0]
	var minSoFar int = nums[0]
	var res int = nums[0]
	for _, w := range nums[1:] {
		maxSoFar, minSoFar = max(max(w, maxSoFar*w), minSoFar*w), min(min(w, minSoFar*w), maxSoFar*w)
		res = max(res, maxSoFar)
	}

	return res
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}
