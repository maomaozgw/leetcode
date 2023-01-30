package p1137

const (
	maxIndex = 38
)

var (
	answers = make([]int, maxIndex)
)

func init() {

	answers[0] = 0
	answers[1] = 1
	answers[2] = 1

	for i := 3; i < maxIndex; i++ {
		answers[i] = answers[i-1] + answers[i-2] + answers[i-3]
	}

}

func tribonacci(n int) int {
	return answers[n]
}
