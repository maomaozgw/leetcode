package p1291

import (
	"sort"
)

var answerList []int

func init() {
	for i := 2; i < 9; i++ {
		var base = 1
		var prev = 1
		for j := 1; j < i; j++ {
			base *= 10
			prev = prev*10 + j + 1
		}
		answerList = append(answerList, prev)
		for j := 2; j <= 10-i; j++ {
			prev = (prev-base*(j-1))*10 + j + i - 1
			answerList = append(answerList, prev)
		}
	}
	answerList = append(answerList, 123456789)
}

func sequentialDigits(low int, high int) []int {
	start := sort.Search(len(answerList), func(i int) bool {
		return answerList[i] >= low
	})
	end := sort.Search(len(answerList), func(i int) bool {
		return answerList[i] > high
	})
	return answerList[start:end]
}
