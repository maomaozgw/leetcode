package p2326

import "github.com/maomaozgw/leetcode/structures/listnode"

type ListNode = listnode.G[int]

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	var result = make([][]int, m)
	var row = make([]int, n)
	for i := range row {
		row[i] = -1
	}
	for i := range result {
		result[i] = make([]int, n)
		copy(result[i], row)
	}
	var left, right, top, bottom = 0, n - 1, 0, m - 1
	for left <= right && top <= bottom && head != nil {
		for i := left; i <= right && head != nil; i++ {
			result[top][i] = head.Val
			head = head.Next
		}
		top++
		if head == nil {
			break
		}
		for j := top; j <= bottom && head != nil; j++ {
			result[j][right] = head.Val
			head = head.Next
		}
		if head == nil {
			break
		}
		right--
		for i := right; i >= left && head != nil; i-- {
			result[bottom][i] = head.Val
			head = head.Next
		}
		if head == nil {
			break
		}
		bottom--
		for i := bottom; i >= top && head != nil; i-- {
			result[i][left] = head.Val
			head = head.Next
		}
		if head == nil {
			break
		}
		left++
	}
	return result
}
