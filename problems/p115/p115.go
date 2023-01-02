package p115

type MinStack struct {
	elements []int
	mins     []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	if len(s.mins) == 0 || val <= s.GetMin() {
		s.mins = append(s.mins, val)
	}
	s.elements = append(s.elements, val)
}

func (s *MinStack) Pop() {
	if s.Top() == s.GetMin() {
		s.mins = s.mins[:len(s.mins)-1]
	}
	s.elements = s.elements[:len(s.elements)-1]
}

func (s *MinStack) Top() int {
	return s.elements[len(s.elements)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}
