package p299

import "fmt"

const (
	start = '0'
)

func getHint(secret string, guess string) string {
	var cow, bull int
	var digitsCnt = make([]int, 10)
	var l = len(secret)
	for i := 0; i < l; i++ {
		digitsCnt[secret[i]-start]++
	}

	for i := 0; i < l; i++ {
		if secret[i] == guess[i] {
			bull++
			digitsCnt[guess[i]-start]--
		} else if digitsCnt[guess[i]-start] > 0 {
			cow++
			digitsCnt[guess[i]-start]--
		}
	}

	for i := 0; i < 10; i++ {
		if digitsCnt[i] < 0 {
			cow += digitsCnt[i]
		}
	}
	return fmt.Sprintf("%dA%dB", bull, cow)

}
