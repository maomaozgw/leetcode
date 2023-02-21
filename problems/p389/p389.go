package p389

const (
	start = 'a'
)

func findTheDifference(s string, t string) byte {
	var cSet = make([]int, 26)
	for i := range s {
		cSet[s[i]-start]++
	}
	for i := range t {
		cSet[t[i]-start]--
		if cSet[t[i]-start] == -1 {
			return t[i]
		}
	}
	return '0'
}
