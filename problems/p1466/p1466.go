package p1466

type edge struct {
	p int
	n int
}

func minReorder(n int, connections [][]int) int {
	var g = map[int][]int{}

	var roads = map[edge]bool{}

	for idx := range connections {
		roads[edge{
			connections[idx][0],
			connections[idx][1],
		}] = true
		g[connections[idx][0]] = append(g[connections[idx][0]], connections[idx][1])
		g[connections[idx][1]] = append(g[connections[idx][1]], connections[idx][0])
	}
	var dfs func(node, parent int)
	var result = 0
	dfs = func(node, parent int) {
		if roads[edge{
			p: parent,
			n: node,
		}] {
			result++
		}
		for _, next := range g[node] {
			if next == parent {
				continue
			}
			dfs(next, node)
		}
	}
	dfs(0, -1)
	return result
}
