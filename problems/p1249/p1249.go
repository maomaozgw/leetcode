package p1249

import "strings"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type Stack struct {
	Current *Node
	Count   int
}

func New() *Stack {
	return &Stack{
		Current: &Node{},
		Count:   0,
	}
}

func (s *Stack) Push(c int) {
	s.Current.Next = &Node{Val: c, Prev: s.Current}
	s.Current = s.Current.Next
	s.Count++
}

func (s *Stack) Pop() (int, bool) {
	if s.Count == 0 {
		return 0, false
	}
	node := s.Current
	s.Current = s.Current.Prev
	node.Prev = nil
	s.Count--
	return node.Val, true
}

func (s *Stack) Peek() int {
	return s.Current.Val
}

func (s *Stack) Len() int {
	return s.Count
}

func minRemoveToMakeValidWithStack(s string) string {
	var left = New()
	var removeIdx = map[int]bool{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			left.Push(i)
		case ')':
			if left.Len() == 0 {
				removeIdx[i] = true
			} else {
				left.Pop()
			}
		default:
			continue
		}
	}
	for {
		idx, ok := left.Pop()
		if !ok {
			break
		}
		removeIdx[idx] = true
	}
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if removeIdx[i] {
			continue
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func minRemoveToMakeValid(s string) string {
	var leftCnt = 0
	var bs = make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(':
			leftCnt++
			bs = append(bs, c)
		case ')':
			if leftCnt == 0 {
				continue
			}
			leftCnt--
			bs = append(bs, c)
		default:
			bs = append(bs, c)
		}
	}
	if len(bs) == leftCnt {
		return ""
	}
	for i := len(bs) - 1; leftCnt > 0; i-- {
		if bs[i] != '(' {
			continue
		}
		bs = append(bs[:i], bs[i+1:]...)
		leftCnt--
	}
	return string(bs)
}
