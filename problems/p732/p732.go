// https://leetcode.com/problems/my-calendar-iii/

package p732

type node struct {
	Next *node
	Val  int
}

func (n *node) Push(v int) *node {
	if n == nil {
		return &node{
			Next: nil,
			Val:  v,
		}
	}
	c := n
	var p *node
	for c != nil {
		if c.Val > v {
			if p == nil {
				p = &node{
					Next: c,
					Val:  v,
				}
				return p
			}
			p.Next = &node{
				Next: c,
				Val:  v,
			}
			return n
		}
		p = c
		c = c.Next
	}
	p.Next = &node{
		Next: nil,
		Val:  v,
	}
	return n
}

type MyCalendarThree struct {
	pointsValue map[int]int
	h           *node
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		pointsValue: map[int]int{},
		h:           nil,
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	if _, ok := this.pointsValue[start]; !ok {
		this.h = this.h.Push(start)
	}
	this.pointsValue[start] += 1
	if _, ok := this.pointsValue[end]; !ok {
		this.h = this.h.Push(end)
	}
	this.pointsValue[end] -= 1
	var result = 0
	var current = 0
	var c = this.h
	for c != nil {
		v := c.Val
		current += this.pointsValue[v]
		result = max(result, current)
		c = c.Next
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
