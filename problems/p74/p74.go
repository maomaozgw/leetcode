// https://leetcode.com/problems/search-a-2d-matrix/description/

package p74

const (
	notFound = -1
)

func searchMatrix(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])
	return binarySearch(matrix, target, 0, m*n-1, n) >= 0
}

func binarySearch(nums [][]int, target int, left, right, n int) int {
	middle := left + (right-left)/2
	if middle == left {
		if nums[middle/n][middle%n] == target {
			return left
		}
		if nums[right/n][right%n] == target {
			return right
		}
		return notFound
	}
	middleVal := nums[middle/n][middle%n]
	if middleVal > target {
		return binarySearch(nums, target, left, middle, n)
	} else if middleVal == target {
		return middle
	} else {
		return binarySearch(nums, target, middle, right, n)
	}
}
