package p1310

func xorQueries(arr []int, queries [][]int) []int {
	var prefixXor = make([]int, len(arr))
	prefixXor[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixXor[i] = prefixXor[i-1] ^ arr[i]
	}
	var result = make([]int, len(queries))
	for i := range queries {
		if queries[i][0] == 0 {
			result[i] = prefixXor[queries[i][1]]
		} else {
			result[i] = prefixXor[queries[i][1]] ^ prefixXor[queries[i][0]-1]
		}
	}
	return result
}
