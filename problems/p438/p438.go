package p438

const (
	start = 'a'
)

func findAnagrams(s string, p string) []int {
	var sourceSet = make([]int, 26)
	for i := 0; i < len(p); i++ {
		sourceSet[p[i]-start]++
	}
	var l = len(p) - 1
	var tmpSet = make([]int, 26)
	var result []int
	var startIdx = 0
	for i := 0; i < len(s); i++ {
		tmpSet[s[i]-start]++
		if i-startIdx != l {
			continue
		}
		if equals(sourceSet, tmpSet) {
			result = append(result, startIdx)
		}
		tmpSet[s[startIdx]-start]--
		startIdx++
	}
	return result
}

func equals(source, target []int) bool {
	for i := 0; i < 26; i++ {
		if source[i] != target[i] {
			return false
		}
	}
	return true
}
