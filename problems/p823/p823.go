// https://leetcode.com/problems/binary-trees-with-factors/

package p823

import (
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	const (
		modVal = 1000000007
	)
	var result = 0
	var valMap = map[int]int{}
	sort.Ints(arr)
	for _, val := range arr {
		valMap[val] = 1
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j <= i; j++ {
			val := arr[i] * arr[j]
			if _, ok := valMap[val]; !ok {
				continue
			}
			if i == j {
				valMap[val] = (valMap[val] + valMap[arr[i]]*valMap[arr[j]]) % modVal
			} else {
				valMap[val] = (valMap[val] + valMap[arr[i]]*valMap[arr[j]]*2) % modVal
			}
		}
	}
	for _, val := range valMap {
		result = (result + val) % modVal
	}
	return result
}
