package p179

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	var numStrs = make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}
	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})
	var result string
	if numStrs[0] == "0" {
		return "0"
	}
	for _, numStr := range numStrs {
		result += numStr
	}
	return result
}
