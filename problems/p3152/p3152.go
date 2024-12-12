package p3152

func isArraySpecial(nums []int, queries [][]int) []bool {
	var prefixSum []int = make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1]
		if nums[i-1]%2 == nums[i]%2 {
			prefixSum[i]++
		}
	}
	var result = make([]bool, len(queries))
	for i := range queries {
		if queries[i][0] == 0 {
			result[i] = prefixSum[queries[i][1]] == 0
		} else {
			result[i] = prefixSum[queries[i][1]]-prefixSum[queries[i][0]] == 0
		}
	}
	return result
}
