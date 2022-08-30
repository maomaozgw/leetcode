// https://leetcode.com/problems/rotate-image/

package p48

func rotate(matrix [][]int) {
	n := len(matrix)
	l := n / 2
	for i := 0; i < l; i++ {
		var (
			x1, y1 = 0 + i, 0 + i
			x2, y2 = x1, n - 1 - i
			x3, y3 = y2, y2
			x4, y4 = x3, y1
		)
		for y1 < n-i-1 {
			matrix[x1][y1], matrix[x4][y4], matrix[x3][y3], matrix[x2][y2] = matrix[x4][y4], matrix[x3][y3], matrix[x2][y2], matrix[x1][y1]
			y1++
			x2++
			y3--
			x4--
		}
	}
}
