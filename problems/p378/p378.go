// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

package p378

import (
	"sort"
)

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	sorted := make([]int, n*n)
	for i := 0; i < n; i++ {
		copy(sorted[(i*n):], matrix[i])
	}
	sort.Ints(sorted)
	return sorted[k-1]
}
