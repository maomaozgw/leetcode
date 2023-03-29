package p2360

func longestCycle(edges []int) int {
	var result = -1

	var visted = make([]int, len(edges))
	var checkStep = 1
	for node := range edges {
		if visted[node] > 0 {
			continue
		}
		startStep := checkStep
		var next = node
		for next != -1 && visted[next] == 0 {
			visted[next] = checkStep
			checkStep++
			next = edges[next]
		}
		if next != -1 && visted[next] >= startStep {
			result = max(result, checkStep-visted[next])
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
