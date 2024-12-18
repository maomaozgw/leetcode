package p1475

import (
	"github.com/emirpasic/gods/v2/stacks/linkedliststack"
)

func finalPrices(prices []int) []int {
	stack := linkedliststack.New[int]()
	for i := len(prices) - 1; i >= 0; i-- {
		for {
			if val, ok := stack.Peek(); !ok {
				break
			} else if val <= prices[i] {
				break
			} else {
				stack.Pop()
			}
		}
		if stack.Empty() {
			stack.Push(prices[i])
			continue
		}
		val, _ := stack.Peek()
		stack.Push(prices[i])
		prices[i] -= val
	}
	return prices
}
