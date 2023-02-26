package p2477

import "math"

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	adjList := make([][]int, n)

	for _, road := range roads {
		u := road[0]
		v := road[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}
	var fuel int64
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {

		rep := 1
		for _, child := range adjList[node] {
			if child != parent {
				rep += dfs(child, node)
			}
		}
		if node != 0 {
			fuel += int64(math.Ceil(float64(rep) / float64(seats)))
		}

		return rep
	}
	dfs(0, -1)
	return fuel
}
