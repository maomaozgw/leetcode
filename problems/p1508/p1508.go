package p1508

import "sort"

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, n*(n+1)/2)
	index := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			sums[index] = sum
			index++
		}
	}
	sort.Ints(sums)
	res := 0
	for i := left - 1; i < right; i++ {
		res += sums[i]
	}
	return res % (1e9 + 7)
}
