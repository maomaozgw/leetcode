package p946

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

func validateStackSequences(pushed []int, popped []int) bool {
	stack := linkedliststack.New()
	var j = 0
	for i := 0; i < len(pushed); i++ {
		stack.Push(pushed[i])
		for !stack.Empty() {
			top, ok := stack.Peek()
			if ok && top.(int) == popped[j] {
				stack.Pop()
				j++
				continue
			}
			break
		}
	}
	return stack.Empty()
}
