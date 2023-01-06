package p1833

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var count = 0
	for ; coins >= 0 && count < len(costs); count++ {
		coins -= costs[count]
	}
	if coins < 0 {
		count--
	}
	return count
}
