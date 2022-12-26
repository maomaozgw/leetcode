package p763

const (
	startC = 'a'
)

func partitionLabels(s string) []int {
	var lastIdx = make([]int, 26)
	for i := 0; i < len(s); i++ {
		lastIdx[s[i]-startC] = i
	}
	var result []int
	for i := 0; i < len(s); i++ {
		last := lastIdx[s[i]-startC]
		for j := i + 1; j <= last; j++ {
			idx := s[j] - startC
			if lastIdx[idx] > last {
				last = lastIdx[idx]
			}
		}
		result = append(result, last-i+1)
		i = last
	}
	return result
}
