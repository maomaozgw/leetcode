package p2825

func canMakeSubsequence(str1 string, str2 string) bool {
	var i, j int
	for i = 0; i < len(str1) && j < len(str2); i++ {
		if str1[i] == str2[j] || str1[i]+1 == str2[j] || str1[i]-25 == str2[j] {
			j++
		}
	}
	return j == len(str2)
}
