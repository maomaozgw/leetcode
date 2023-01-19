package p509

var (
	answers = make([]int, 31)
)

func init() {
	answers[0] = 0
	answers[1] = 1
	for i := 2; i < 31; i++ {
		answers[i] = answers[i-1] + answers[i-2]
	}
}

func fib(n int) int {
	return answers[n]
}
