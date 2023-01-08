package p215

func findKthLargest(nums []int, k int) int {
	var cnt = make([]int, 20001)
	for i := range nums {
		cnt[nums[i]+10000]++
	}
	for i := 20000; i > -1; i-- {
		k -= cnt[i]
		if k <= 0 {
			return i - 10000
		}
	}
	return cnt[0] - 10000
}
