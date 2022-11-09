// https://leetcode.com/problems/online-stock-span/

package p901

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

type StockSpanner struct {
	values *linkedliststack.Stack
}

func Constructor() StockSpanner {
	s := StockSpanner{
		values: linkedliststack.New(),
	}
	s.values.Push([]int{1<<31 - 1, 0})
	return s
}

func (this *StockSpanner) Next(price int) int {

	last, _ := this.values.Peek()
	lastVal := last.([]int)
	idx := lastVal[1] + 1
	if price < lastVal[0] {
		this.values.Push([]int{price, idx})
		return 1
	}
	for {
		this.values.Pop()
		last, _ = this.values.Peek()
		lastVal = last.([]int)
		if price < lastVal[0] {
			this.values.Push([]int{price, idx})
			return idx - lastVal[1]
		}
	}
}
