package p560

func subarraySum(nums []int, k int) int {
	var result = 0
	var n = len(nums)
	var preSum = make([]int, n)
	preSum[0] = nums[0]
	// 记录与 k 差 x 的个数
	var cache = map[int]int{}
	cache[0] = 1
	// 1, 1, 1
	for i := 1; i < n; i++ {
		// 1, 2, 3
		preSum[i] = nums[i] + preSum[i-1]
	}
	// k = 2
	for i := 0; i < n; i++ {
		val := preSum[i]
		if cnt, ok := cache[val-k]; ok {
			result += cnt
		}
		cache[val] += 1
	}
	return result
}
