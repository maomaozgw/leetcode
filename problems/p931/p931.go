// https://leetcode.com/problems/minimum-falling-path-sum/description/

package p931

const (
	intMax = 1 << 31
)

func minFallingPathSum(matrix [][]int) int {
	var dp = map[int]int{}
	var m = len(matrix)
	var result = intMax
	for i := 0; i < m; i++ {
		result = min(result, getSum(matrix, m, 0, i, dp))
	}
	return result
}

func getSum(matrix [][]int, m, i, j int, cache map[int]int) int {
	if i == m || j == m || j < 0 {
		return intMax
	}
	idx := i*1000 + j
	if val, ok := cache[idx]; ok {
		return val
	}
	current := matrix[i][j]
	if i == m-1 {
		return current
	}
	min1 := getSum(matrix, m, i+1, j, cache)
	min2 := getSum(matrix, m, i+1, j-1, cache)
	min3 := getSum(matrix, m, i+1, j+1, cache)
	current += min(min(min1, min2), min3)
	cache[idx] = current
	return current
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
