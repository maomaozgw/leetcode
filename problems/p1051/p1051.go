package p1051

import "sort"

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	var result = 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			result++
		}
	}
	return result
}
