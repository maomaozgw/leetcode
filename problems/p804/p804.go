package p804

import "strings"

const (
	a = byte('a')
)

var (
	morseInfo = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
)

func uniqueMorseRepresentations(words []string) int {
	resultMap := map[string]bool{}
	for i := 0; i < len(words); i++ {
		resultMap[transformOne(words[i])] = true
	}
	return len(resultMap)
}

func transformOne(s string) string {
	b := strings.Builder{}
	for i := 0; i < len(s); i++ {
		b.WriteString(morseInfo[s[i]-a])
	}
	return b.String()
}
