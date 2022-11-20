package p224

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

const (
	opInit = iota
	opNone
	opAdd
	opDelete
)

type item struct {
	val int
	op  int
}

func (i *item) calc(nextNum int) {
	switch i.op {
	case opInit:
		i.val = nextNum
		i.op = opNone
	case opAdd:
		i.val += nextNum
		i.op = opNone
	case opDelete:
		i.val -= nextNum
		i.op = opNone
	case opNone:
	default:
		panic("unexpected val")
	}
}

func calculate(s string) int {
	stack := linkedliststack.New()
	var last *item = &item{
		val: 0,
		op:  opInit,
	}
	var currentVal int
	for i := 0; i < len(s); i++ {
		element := s[i]
		switch element {
		case '+', '-':
			last.calc(currentVal)
			currentVal = 0
			if element == '+' {
				last.op = opAdd
			} else if element == '-' {
				last.op = opDelete
			}
		case ' ':
			continue
		case '(':
			stack.Push(last)
			last = &item{}
		case ')':
			last.calc(currentVal)
			currentVal = 0
			prev, ok := stack.Pop()
			if ok {
				preItem := prev.(*item)
				preItem.calc(last.val)
				last = preItem
			}
		default:
			currentVal = currentVal*10 + int(element-'0')
		}
	}
	last.calc(currentVal)
	for {
		prev, ok := stack.Pop()
		if !ok {
			break
		}
		preItem := prev.(*item)
		preItem.calc(last.val)

		last = preItem
	}
	return last.val
}
