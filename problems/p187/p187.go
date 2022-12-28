package p187

func findRepeatedDnaSequences(s string) []string {
	l := len(s) - 10
	var seqMap = map[string]int{}
	var result []string
	for i := 0; i <= l; i++ {
		subStr := s[i : i+10]
		if seqMap[subStr] == 1 {
			result = append(result, subStr)
		}
		seqMap[subStr]++
	}
	return result
}
