package p2409

func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	for i := 1; i < len(sentence); i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}
