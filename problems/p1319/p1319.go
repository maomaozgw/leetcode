package p1319

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		// cannot connection
		return -1
	}
	var g = map[int]map[int]bool{}
	for idx := range connections {
		if _, ok := g[connections[idx][0]]; !ok {
			g[connections[idx][0]] = map[int]bool{}
		}
		g[connections[idx][0]][connections[idx][1]] = true
		if _, ok := g[connections[idx][1]]; !ok {
			g[connections[idx][1]] = map[int]bool{}
		}
		g[connections[idx][1]][connections[idx][0]] = true
	}

	var visited = make([]bool, n)
	var dfs func(node int) int

	dfs = func(node int) int {
		if visited[node] {
			return 0
		}
		visited[node] = true
		for peer := range g[node] {
			dfs(peer)
		}
		return 1
	}
	var result = 0
	for i := 0; i < n; i++ {
		result += dfs(i)
	}
	return result - 1
}
