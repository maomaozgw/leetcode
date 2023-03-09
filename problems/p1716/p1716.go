package p1716

func totalMoney(n int) int {
	var cnt7 = n / 7
	var remain = n % 7
	var sum = (1 + cnt7*2 + remain) * remain / 2
	if cnt7 > 0 {
		sum += cnt7*28 + (cnt7-1)*(cnt7)/2*7
	}
	return sum
}
