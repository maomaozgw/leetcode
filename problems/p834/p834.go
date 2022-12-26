package p834

func sumOfDistancesInTree(n int, edges [][]int) []int {
	var nodeDict = map[int][]int{}
	for i := range edges {
		edge := edges[i]
		start, end := edge[0], edge[1]
		if _, ok := nodeDict[start]; !ok {
			nodeDict[start] = []int{}
		}
		if _, ok := nodeDict[end]; !ok {
			nodeDict[end] = []int{}
		}
		nodeDict[start] = append(nodeDict[start], end)
		nodeDict[end] = append(nodeDict[end], start)
	}
	res, count := make([]int, n), make([]int, n)

	var postDfs func(node, parent int)
	postDfs = func(node, parent int) {
		for _, dest := range nodeDict[node] {
			if dest != parent {
				postDfs(dest, node)
				res[node] += res[dest] + count[dest]
				count[node] += count[dest]
			}
		}
		count[node]++
	}

	var preDfs func(node, parent int)
	preDfs = func(node, parent int) {
		for _, dest := range nodeDict[node] {
			if dest != parent {
				res[dest] = res[node] - count[dest] + n - count[dest]
				preDfs(dest, node)
			}
		}
	}

	postDfs(0, -1)
	preDfs(0, -1)

	return res
}
