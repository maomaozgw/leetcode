package p520

const (
	A = 'A'
	Z = 'Z'
	a = 'a'
	z = 'z'
)

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return true
	}
	var hasLower = false
	var hasUpper = false
	if a <= word[0] && word[0] <= z {
		hasLower = true
	}
	for i := 1; i < len(word); i++ {
		if a <= word[i] && word[i] <= z {
			hasLower = true
		} else if A <= word[i] && word[i] <= Z {
			hasUpper = true
		} else {
			panic("invalid input")
		}
		if hasUpper && hasLower {
			return false
		}
	}
	return true
}
