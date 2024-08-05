package p2053

func kthDistinct(arr []string, k int) string {
	distinctMap := make(map[string]int)
	for _, s := range arr {
		distinctMap[s]++
	}
	distinctArr := make([]string, 0)
	for _, s := range arr {
		if distinctMap[s] == 1 {
			distinctArr = append(distinctArr, s)
		}
	}
	if k > len(distinctArr) {
		return ""
	}
	return distinctArr[k-1]
}
