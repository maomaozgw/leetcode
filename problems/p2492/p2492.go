package p2492

import (
	"math"
)

type node struct {
	n int
	s int
}

func minScore(n int, roads [][]int) int {
	var result = math.MaxInt

	var nodeMap = map[int][]node{}
	for _, edge := range roads {
		nodeMap[edge[0]] = append(nodeMap[edge[0]], node{
			n: edge[1],
			s: edge[2],
		})
		nodeMap[edge[1]] = append(nodeMap[edge[1]], node{
			n: edge[0],
			s: edge[2],
		})
	}
	var visted = make([]int, n+1)

	var q []int
	q = append(q, 1)
	visted[1] = 1
	for len(q) > 0 {
		var newQ []int
		for _, n := range q {
			for _, node := range nodeMap[n] {
				result = min(result, node.s)
				if visted[node.n] == 1 {
					continue
				}
				visted[node.n] = 1
				newQ = append(newQ, node.n)
			}
		}
		q = newQ
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
