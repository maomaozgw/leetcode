package p974

func subarraysDivByK(nums []int, k int) int {
	d := make(map[int]int)
	pres, ans := 0, 0
	d[0] = 1
	for _, n := range nums {
		pres += n
		// for handling negative remainder
		// python is intelligent, it internally do this thing
		pres = (pres%k + k) % k

		ans += d[pres]
		d[pres]++
	}
	return ans
}

func subarraysDivByKTooLong(nums []int, k int) int {

	var result = 0
	var n = len(nums)
	var preSum = make([]int, n+1)
	preSum[0] = nums[0]
	for i := 1; i < n+1; i++ {
		// 1, 2, 3
		preSum[i] = nums[i-1] + preSum[i-1]
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if (preSum[j+1]-preSum[i])%k == 0 {
				result++
			}
		}
	}
	return result

}
