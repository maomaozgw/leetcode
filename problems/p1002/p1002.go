package p1002

func commonChars(words []string) []string {
	res := make([]string, 0)
	count := make([]int, 26)
	for i := 0; i < 26; i++ {
		count[i] = 100
	}
	for _, word := range words {
		temp := make([]int, 26)
		for _, c := range word {
			temp[c-'a']++
		}
		for i := 0; i < 26; i++ {
			if temp[i] < count[i] {
				count[i] = temp[i]
			}
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < count[i]; j++ {
			res = append(res, string(byte(i)+'a'))
		}
	}
	return res
}
