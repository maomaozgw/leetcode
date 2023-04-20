package p1337

import "sort"

type rowCounter struct {
	r int
	c int
}

func kWeakestRows(mat [][]int, k int) []int {
	var rows []rowCounter
	for i := range mat {
		c := binarySearch(mat[i])
		rows = append(rows, rowCounter{
			r: i,
			c: c,
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].c == rows[j].c {
			return rows[i].r < rows[j].r
		}
		return rows[i].c < rows[j].c
	})
	var result = make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = rows[i].r
	}
	return result
}

func binarySearch(nums []int) int {
	var left, right = 0, len(nums)

	for left < right {
		var mid = (right + left) >> 1
		if nums[mid] == 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
