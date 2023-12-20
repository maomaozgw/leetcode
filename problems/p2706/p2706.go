package p2706

func buyChoco(prices []int, money int) int {
	var min0, min1 = 10000, 10000
	for i := range prices {
		if prices[i] < min0 {
			min1 = min0
			min0 = prices[i]
		} else if prices[i] < min1 {
			min1 = prices[i]
		}
	}
	remain := money - min0 - min1
	if remain < 0 {
		return money
	}
	return remain
}
