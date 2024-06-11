package p1122

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	relativeMap := map[int]int{}
	for idx, val := range arr2 {
		relativeMap[val] = idx
	}
	sort.Slice(arr1, func(i, j int) bool {
		iVal, iExist := relativeMap[arr1[i]]
		jVal, jExist := relativeMap[arr1[j]]
		if iExist && jExist {
			return iVal < jVal
		}
		if iExist {
			return true
		}
		if jExist {
			return false
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
