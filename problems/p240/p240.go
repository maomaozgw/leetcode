// https://leetcode.com/problems/search-a-2d-matrix-ii/

package p240

func searchMatrix(matrix [][]int, target int) bool {
	for i := range matrix {
		n := len(matrix[i])
		left, right := matrix[i][0], matrix[i][n-1]
		if target == left || target == right {
			return true
		}
		if target < left || target > right {
			continue
		}
		if binarySearch(matrix[i], target, 0, n-1) {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target, left, right int) bool {
	middle := left + (right-left)/2
	if middle == left {
		if nums[middle] == target {
			return true
		}
		if nums[right] == target {
			return true
		}
		return false
	}
	if nums[middle] > target {
		return binarySearch(nums, target, left, middle)
	} else if nums[middle] == target {
		return true
	} else {
		return binarySearch(nums, target, middle, right)
	}
}
