// https://leetcode.com/problems/make-the-string-great/

package p1544

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

const (
	offsetVal  = 'A' - 'a'
	offsetVal2 = -offsetVal
)

func makeGood(s string) string {
	var stack = linkedliststack.New()
	for i := 0; i < len(s); i++ {
		stack.Push(s[i])
		for i < len(s)-1 {
			last, ok := stack.Peek()
			if !ok {
				break
			}
			var offset = int(last.(byte)) - int(s[i+1])
			if offset == offsetVal || offset == offsetVal2 {
				stack.Pop()
				i++
				continue
			}
			break
		}
	}
	var ans = make([]byte, stack.Size())
	var values = stack.Values()
	for i := 0; i < len(values); i++ {
		ans[len(values)-1-i] = values[i].(byte)
	}
	return string(ans)
}
