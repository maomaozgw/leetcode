package p1823

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func findTheWinnerLinkedList(n int, k int) int {
	start := &Node{
		Val: 1,
	}
	// build cycle
	c := start
	for i := 2; i < n+1; i++ {
		c.Next = &Node{Val: i, Prev: c}
		c = c.Next
	}
	c.Next = start
	start.Prev = c
	c = start
	for c.Next != c {
		for i := 1; i < k; i++ {
			c = c.Next
		}
		c.Prev.Next = c.Next
		c.Next.Prev = c.Prev
		tmp := c
		c = c.Next
		tmp.Prev, tmp.Next = nil, nil
	}
	return c.Val
}

func findTheWinner(n, k int) int {
	var l = make([]int, n)
	for i := 1; i <= n; i++ {
		l[i-1] = i
	}
	var idx int
	for len(l) > 1 {
		idx = (idx - 1 + k) % len(l)
		l = append(l[:idx], l[idx+1:]...)
	}
	return l[0]
}
