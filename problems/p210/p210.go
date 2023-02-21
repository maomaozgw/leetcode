package p210

func findOrder(numCourses int, prerequisites [][]int) []int {

	var graph = map[int][]int{}
	var inDegree = make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		inDegree[pre[0]]++
	}

	var q []int
	for i, val := range inDegree {
		if val == 0 {
			q = append(q, i)
		}
	}

	var res []int

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		res = append(res, curr)
		for _, depend := range graph[curr] {
			inDegree[depend]--
			if inDegree[depend] == 0 {
				q = append(q, depend)
			}
		}
	}
	for _, val := range inDegree {
		if val != 0 {
			return []int{}
		}
	}
	return res
}
