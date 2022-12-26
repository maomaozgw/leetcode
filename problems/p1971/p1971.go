// https://leetcode.com/problems/find-if-path-exists-in-graph/description/

package p1971

func validPath(n int, edges [][]int, source int, destination int) bool {
	var walkedPath = map[int]bool{}
	var graphInfo = map[int][]int{}
	for _, item := range edges {
		start, end := item[0], item[1]
		if _, ok := graphInfo[start]; !ok {
			graphInfo[start] = []int{end}
		} else {
			graphInfo[start] = append(graphInfo[start], end)
		}

		if _, ok := graphInfo[end]; !ok {
			graphInfo[end] = []int{start}
		} else {
			graphInfo[end] = append(graphInfo[end], start)
		}
	}
	q := graphInfo[source]
	for len(q) > 0 {
		var newQ []int
		for _, val := range q {
			if val == destination {
				return true
			}
			if walkedPath[val] {
				continue
			}
			walkedPath[val] = true
			if _, ok := graphInfo[val]; ok {
				newQ = append(newQ, graphInfo[val]...)
			}
		}
		q = newQ
	}
	return false
}
