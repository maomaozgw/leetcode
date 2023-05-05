package p1456

const (
	startA = 'a'
)

var (
	vowelSet []bool
)

func init() {
	vowelSet = make([]bool, 26)
	vowelSet['a'-startA] = true
	vowelSet['e'-startA] = true
	vowelSet['i'-startA] = true
	vowelSet['o'-startA] = true
	vowelSet['u'-startA] = true
}

func maxVowels(s string, k int) int {
	var result int
	var current = 0
	for i := 0; i < k; i++ {
		if vowelSet[s[i]-startA] {
			current++
		}
	}
	result = current
	for i := k; i < len(s); i++ {
		if vowelSet[s[i]-startA] {
			current++
		}
		if vowelSet[s[i-k]-startA] {
			current--
		}
		result = max(result, current)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
