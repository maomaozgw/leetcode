package listnode

import (
	"fmt"
	"strings"
)

func (g *G[T]) String() string {
	c := g
	var valueList []string
	for c != nil {
		valueList = append(valueList, fmt.Sprintf("%v", c.V))
	}
	return strings.Join(valueList, "->")
}

func (g *G[T]) Value() T {
	return g.V
}

func (g *G[T]) Next() *G[T] {
	return g.N
}

func (g *D[T]) String() string {
	c := g
	var valueList []string
	for c != nil {
		valueList = append(valueList, fmt.Sprintf("%v", c.V))
	}
	return strings.Join(valueList, "->")
}

func (g *D[T]) Value() T {
	return g.V
}

func (g *D[T]) Next() *D[T] {
	return g.N
}

func Print[T comparable](head *G[T]) {
	for head != nil {
		fmt.Print(head.V, " ")
		head = head.N
	}
	fmt.Print("\n")
}

func DeepEqual[T comparable](one *G[T], another *G[T]) bool {
	var (
		p1 = one
		p2 = another
	)

	for {
		if p1 == nil && p2 == nil {
			return true
		} else if p1 == nil || p2 == nil {

			return false
		}
		if p1.Value() != p2.Value() {
			return false
		}
		p1 = p1.Next()
		p2 = p2.Next()
	}
}
