package p2657

func findThePrefixCommonArray(A []int, B []int) []int {
	numMap := make([]int, len(A)+1)
	var result = make([]int, len(A))
	var comm = 0
	for i := 0; i < len(A); i++ {
		numMap[A[i]]++
		if numMap[A[i]] == 2 {
			comm++
		}
		numMap[B[i]]++
		if numMap[B[i]] == 2 {
			comm++
		}
		result[i] = comm
	}
	return result
}
