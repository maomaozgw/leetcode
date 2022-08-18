package p20

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

func isValid(s string) bool {

	l := linkedliststack.New()
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isOpen(c) {
			l.Push(c)
		} else {
			be, _ := l.Peek()
			if be == nil {
				return false
			}
			b := be.(byte)
			if !canClose(c, b) {
				return false
			}
			_, _ = l.Pop()
		}
	}
	return l.Empty()
}

func isOpen(c byte) bool {
	return c == '(' || c == '{' || c == '['
}

func canClose(c byte, b byte) bool {
	return (c == ')' && b == '(') || (c == '}' && b == '{') || (c == ']' && b == '[')
}
