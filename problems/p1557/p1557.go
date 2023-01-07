package p1557

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	for _, edge := range edges {
		a := find(edge[0], parents)
		b := find(edge[1], parents)
		if a != b && parents[edge[1]] == edge[1] {
			parents[b] = a
		}
	}
	var res []int
	for i := range parents {
		if i == parents[i] {
			res = append(res, i)
		}
	}
	return res
}

func find(x int, parents []int) int {
	for x != parents[x] {
		p := parents[x]
		x, parents[x] = p, parents[p]
	}
	return x
}
