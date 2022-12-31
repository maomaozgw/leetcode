// https://leetcode.com/problems/sum-of-subarray-minimums/

package p907

import "fmt"

// https://zhuanlan.zhihu.com/p/371544831
// Mono Stack

func sumSubarrayMins(arr []int) int {
	arr = append(arr, 0)
	stack, res, j := []int{-1}, 0, 0
	for i, n := range arr {
		fmt.Printf("Loop %d: n = %d", i, n)
		for len(stack) > 1 && arr[stack[len(stack)-1]] > n {
			j, stack = stack[len(stack)-1], stack[:len(stack)-1]
			res += (j - stack[len(stack)-1]) * (i - j) * arr[j]
			fmt.Printf(", j = %d(%d), res = %d", j, arr[j], res)
			res %= 1_000_000_007
		}
		stack = append(stack, i)
		fmt.Printf(",stack = %v\n", stack)
	}
	return res
}
