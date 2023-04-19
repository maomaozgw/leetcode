// https://leetcode.com/problems/search-a-2d-matrix/description/

package p74

func searchMatrix(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])
	var left, right = 0, m * n
	for left < right {
		var mid = (left + right) >> 1
		var val = matrix[mid/n][mid%n]
		if val == target {
			return true
		} else if val > target {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return false
}
